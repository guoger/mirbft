/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mirbft

import (
	pb "github.com/IBM/mirbft/mirbftpb"

	"go.uber.org/zap"
)

type SequenceState int

const (
	Uninitialized SequenceState = iota
	Preprepared
	Digested
	InvalidBatch
	Validated
	Prepared
	Committed
)

type sequence struct {
	myConfig *Config

	epochConfig *epochConfig

	state SequenceState

	// Entry's Batch field is unset until after state >= Preprepared
	entry *Entry

	// Digest is not set until after state >= Digested
	digest []byte

	prepares map[string]map[NodeID]struct{}
	commits  map[string]map[NodeID]struct{}
}

func newSequence(epochConfig *epochConfig, myConfig *Config, entry *Entry) *sequence {
	return &sequence{
		myConfig:    myConfig,
		epochConfig: epochConfig,
		entry:       entry,
		state:       Uninitialized,
		prepares:    map[string]map[NodeID]struct{}{},
		commits:     map[string]map[NodeID]struct{}{},
	}
}

// applyPreprepare attempts to apply a batch from a preprepare message to the state machine.
// If the state machine is not in the Uninitialized state, it returns an error.  Otherwise,
// It transitions to Preprepared and returns a ValidationRequest message.
func (s *sequence) applyPreprepareMsg(batch [][]byte) *Actions {
	if s.state != Uninitialized {
		s.myConfig.Logger.Panic("illegal state for preprepare", zap.Uint64(SeqNoLog, s.entry.SeqNo), zap.Uint64(EpochLog, s.epochConfig.number), zap.Int("CurrentState", int(s.state)), zap.Int("Expected", int(Uninitialized)))
	}

	s.state = Preprepared
	s.entry.Batch = batch

	return &Actions{
		Digest: []*Entry{s.entry},
	}
}

func (s *sequence) applyDigestResult(digest []byte) *Actions {
	if s.state != Preprepared {
		s.myConfig.Logger.Panic("illegal state for digest result", zap.Uint64(SeqNoLog, s.entry.SeqNo), zap.Uint64(EpochLog, s.epochConfig.number), zap.Int("CurrentState", int(s.state)), zap.Int("Expected", int(Preprepared)))
	}

	s.state = Digested
	s.digest = digest

	return &Actions{
		Validate: []*Entry{s.entry},
	}
}

func (s *sequence) applyValidateResult(valid bool) *Actions {
	if s.state != Digested {
		s.myConfig.Logger.Panic("illegal state for validate result", zap.Uint64(SeqNoLog, s.entry.SeqNo), zap.Uint64(EpochLog, s.epochConfig.number), zap.Int("CurrentState", int(s.state)), zap.Int("Expected", int(Digested)))
	}

	if !valid {
		s.state = InvalidBatch
		// TODO return a view change / suspect message
		return &Actions{}
	}

	s.state = Validated

	return &Actions{
		Broadcast: []*pb.Msg{
			{
				Type: &pb.Msg_Prepare{
					Prepare: &pb.Prepare{
						SeqNo:  s.entry.SeqNo,
						Epoch:  s.entry.Epoch,
						Digest: s.digest,
					},
				},
			},
		},
	}
}

func (s *sequence) applyPrepareMsg(source NodeID, digest []byte) *Actions {
	// TODO, if the digest is known, mark a mismatch as oddity
	agreements := s.prepares[string(digest)]
	if agreements == nil {
		agreements = map[NodeID]struct{}{}
		s.prepares[string(digest)] = agreements
	}
	agreements[source] = struct{}{}

	if s.state != Validated {
		return &Actions{}
	}

	// Do not prepare unless we have sent our prepare as well
	if _, ok := agreements[NodeID(s.myConfig.ID)]; !ok {
		return &Actions{}
	}

	// We do require 2f+1 prepares (instead of 2f), a prepare is implicitly added for the leader
	requiredPrepares := s.epochConfig.intersectionQuorum()

	if len(agreements) < requiredPrepares {
		return &Actions{}
	}

	s.state = Prepared

	return &Actions{
		Broadcast: []*pb.Msg{
			{
				Type: &pb.Msg_Commit{
					Commit: &pb.Commit{
						SeqNo:  s.entry.SeqNo,
						Epoch:  s.entry.Epoch,
						Digest: s.digest,
					},
				},
			},
		},
	}
}

func (s *sequence) applyCommitMsg(source NodeID, digest []byte) *Actions {
	// TODO, if the digest is known, mark a mismatch as oddity
	agreements := s.commits[string(digest)]
	if agreements == nil {
		agreements = map[NodeID]struct{}{}
		s.commits[string(digest)] = agreements
	}
	agreements[source] = struct{}{}

	if s.state != Prepared {
		return &Actions{}
	}

	// Do not commit unless we have sent a commit
	if _, ok := agreements[NodeID(s.myConfig.ID)]; !ok {
		return &Actions{}
	}

	requiredCommits := s.epochConfig.intersectionQuorum()

	if len(agreements) < requiredCommits {
		return &Actions{}
	}

	s.state = Committed

	return &Actions{
		Commit: []*Entry{s.entry},
	}
}

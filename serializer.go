/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mirbft

import (
	"sync"

	pb "github.com/IBM/mirbft/mirbftpb"

	"github.com/pkg/errors"
	// "go.uber.org/zap"
)

type step struct {
	Source uint64
	Msg    *pb.Msg
}

type clientReq struct {
	clientID []byte
	replyC   chan *clientWaiter
}

// serializer provides a single threaded way to access the Mir state machine
// and passes work to/from the state machine.
type serializer struct {
	actionsC chan Actions
	doneC    <-chan struct{}
	clientC  chan *clientReq
	propC    chan *pb.RequestData
	resultsC chan ActionResults
	statusC  chan chan<- *Status
	stepC    chan step
	tickC    chan struct{}
	errC     chan struct{}

	exitMutex    sync.Mutex
	exitErr      error
	exitStatus   *Status
	stateMachine *stateMachine
}

func newSerializer(stateMachine *stateMachine, doneC <-chan struct{}) *serializer {
	s := &serializer{
		actionsC:     make(chan Actions),
		doneC:        doneC,
		propC:        make(chan *pb.RequestData),
		clientC:      make(chan *clientReq),
		resultsC:     make(chan ActionResults),
		statusC:      make(chan chan<- *Status),
		stepC:        make(chan step),
		tickC:        make(chan struct{}),
		errC:         make(chan struct{}),
		stateMachine: stateMachine,
	}
	go s.run()
	return s
}

func (s *serializer) getExitErr() error {
	s.exitMutex.Lock()
	defer s.exitMutex.Unlock()
	return s.exitErr
}

// run must be single threaded and is therefore hidden to prevent accidental capture
// of other go routines.
func (s *serializer) run() {
	defer func() {
		s.exitMutex.Lock()
		defer s.exitMutex.Unlock()
		close(s.errC)
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				s.exitErr = errors.Wrapf(err, "serializer caught panic")
			} else {
				s.exitErr = errors.Errorf("panic in statemachine: %v", r)
			}
			panic(r)
		} else {
			s.exitErr = ErrStopped
		}
		s.exitStatus = s.stateMachine.status()
	}()

	actions := &Actions{}
	var actionsC chan<- Actions
	for {
		select {
		case data := <-s.propC:
			// s.stateMachine.myConfig.Logger.Debug("serializer receiving", zap.String("type", "proposal"))
			actions.Append(s.stateMachine.propose(data))
		case req := <-s.clientC:
			req.replyC <- s.stateMachine.clientWaiter(req.clientID)
		case step := <-s.stepC:
			// s.stateMachine.myConfig.Logger.Debug("serializer receiving", zap.String("type", "step"))
			actions.Append(s.stateMachine.step(NodeID(step.Source), step.Msg))
		case actionsC <- *actions:
			// s.stateMachine.myConfig.Logger.Debug("serializer sent actions")
			actions.Clear()
			actionsC = nil
			continue
		case results := <-s.resultsC:
			// s.stateMachine.myConfig.Logger.Debug("serializer receiving", zap.String("type", "results"))
			actions.Append(s.stateMachine.processResults(results))
		case statusReq := <-s.statusC:
			// s.stateMachine.myConfig.Logger.Debug("serializer receiving", zap.String("type", "status"))
			select {
			case statusReq <- s.stateMachine.status():
			case <-s.doneC:
			}
		case <-s.tickC:
			// s.stateMachine.myConfig.Logger.Debug("serializer receiving", zap.String("type", "tick"))
			actions.Append(s.stateMachine.tick())
		case <-s.doneC:
			// s.stateMachine.myConfig.Logger.Debug("serializer receiving", zap.String("type", "done"))
			return
		}

		// We unconditionally re-enable the actions channel after any event is injected into the system
		// which will mean some zero-length actions get sent to the consumer.  This isn't optimal,
		// but, I've convinced myself that's okay for a couple reasons:
		// 1) Under stress, processing the actions will take enough time for the actions channel to
		// become available again anyway.
		// 2) For tests and visualizations, it's very nice being able to guarantee that we have
		// the latest set of actions (in the other model, it's unclear whether a call to the actions
		// channel will ever unblock.
		actionsC = s.actionsC
	}
}

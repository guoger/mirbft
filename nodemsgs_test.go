/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mirbft

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"

	pb "github.com/IBM/mirbft/mirbftpb"
)

var _ = FDescribe("NodeMsg", func() {
	var (
		nodeID        NodeID
		networkConfig *pb.NetworkConfig
		myConfig      *Config
		clientWindows *clientWindows
		oddities      *oddities
	)

	BeforeEach(func() {
		nodeID = NodeID(0)
		networkConfig = &pb.NetworkConfig{
			CheckpointInterval: 2,
			F:                  0,
			Nodes:              []uint64{0},
			NumberOfBuckets:    1,
			MaxEpochLength:     10,
		}
		myConfig = &Config{
			ID:     uint64(0),
			Logger: zap.NewExample(),
			BatchParameters: BatchParameters{
				CutSizeBytes: 1,
			},
			SuspectTicks:         4,
			NewEpochTimeoutTicks: 8,
			BufferSize:           500,
		}
		clientWindows = nil
		oddities = nil
	})

	It("creates nodemsg", func() {
		n := newNodeMsgs(nodeID, networkConfig, myConfig, clientWindows, oddities)
		Expect(n).NotTo(BeNil())
	})
})

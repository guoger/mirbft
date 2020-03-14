package main

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"hash"
	"os"
	"time"

	"github.com/IBM/mirbft"
	"github.com/IBM/mirbft/example/network"
	pb "github.com/IBM/mirbft/mirbftpb"
	"github.com/IBM/mirbft/sample"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	logger := zap.NewExample()

	f := os.Args[1]
	config, err := network.LoadConfig(f)
	check(err)

	// Create mir instance
	networkConfig := mirbft.StandardInitialNetworkConfig(len(config.Peers))

	mirConfig := &mirbft.Config{
		ID:                   config.ID,
		Logger:               logger,
		BatchParameters:      mirbft.BatchParameters{CutSizeBytes: 1},
		HeartbeatTicks:       1,
		SuspectTicks:         100,
		NewEpochTimeoutTicks: 400,
	}

	doneC := make(chan struct{})
	defer close(doneC)

	node, err := mirbft.StartNewNode(mirConfig, doneC, networkConfig)
	check(err)

	// Create transport
	t, err := network.NewTransport(logger, config)
	check(err)

	t.Handle(func(id uint64, data []byte) {
		msg := &pb.Msg{}
		err := proto.Unmarshal(data, msg)
		check(err)

		err = node.Step(context.TODO(), id, msg)
		if err != nil {
			logger.Warn(fmt.Sprintf("Failed to step message to mir node: %s", err))
		}
	})

	err = t.Start()
	check(err)
	defer t.Close()

	commitC := make(chan *pb.QEntry, 1000)
	defer close(commitC)

	processor := &sample.SerialProcessor{
		Link: t,
		Validator: sample.ValidatorFunc(func(result *mirbft.Request) error {
			if result.Source != binary.LittleEndian.Uint64(result.ClientRequest.ClientId) {
				return fmt.Errorf("mis-matched originating replica and client id")
			}
			return nil
		}),
		Hasher: func() hash.Hash { return md5.New() },
		Committer: &sample.SerialCommitter{
			Log: &network.FakeLog{
				CommitC: commitC,
			},
			OutstandingSeqNos:      map[uint64]*mirbft.Commit{},
			OutstandingCheckpoints: map[uint64]struct{}{},
		},
		Node: node,
	}

	go func() {
		ticker := time.NewTicker(1000 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				node.Tick()
			case actions := <-node.Ready():
				results := processor.Process(&actions)
				node.AddResults(*results)
			case <-node.Err():
				_, err := node.Status(context.Background())
				logger.Warn(err.Error())
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case entry := <-commitC:
				for _, req := range entry.Requests {
					fmt.Printf("### Committing ReqNo: %d\n", req.ReqNo)
				}
			case <-doneC:
				return
			}
		}
	}()

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	err := node.Propose(context.TODO(), []byte(r.URL.Path[1:]))
	//	if err != nil {
	//		w.Write([]byte(err.Error()))
	//	}
	//	w.Write([]byte("Success"))
	//})
	//http.ListenAndServe(":8080", nil)

	if config.ID == 0 {
		for i := 1; i < 10; i++ {
			fmt.Printf("Proposing %d\n", i)
			req := &pb.RequestData{
				ClientId:  []byte("client-1"),
				ReqNo:     uint64(i),
				Data:      []byte("data"),
				Signature: []byte("signature"),
			}
			err := node.Propose(context.TODO(), false, req)
			check(err)
			fmt.Printf("Proposed %d\n", i)
			time.Sleep(10 * time.Second)
		}
	} else {
		time.Sleep(time.Hour)
	}

}
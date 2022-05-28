package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/tonussi/studygo/pkg/communication"
	hashicorpraft "github.com/tonussi/studygo/pkg/ordering/hashicorp-raft"
	"github.com/tonussi/studygo/pkg/proxy"
)

var (
	listenAddr     = flag.String("l", "localhost:8000", "listen requests address")
	deliveryAddr   = flag.String("d", "localhost:8001", "delivery server address")
	listenJoinAddr = flag.String("k", "localhost:9000", "listen join requests address")
	joinAddr       = flag.String("j", "localhost:9000", "join address")
)

func main() {
	flag.Parse()

	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		log.Fatal("node id must be set")
	}

	raftAddr := os.Getenv("PROTOCOL_IP") + ":" + os.Getenv("PROTOCOL_PORT")

	httpCommunicator, err := communication.NewHTTPCommunicator(
		*listenAddr,
		*deliveryAddr,
		5,
		2*time.Second,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	hashicoprRaftOrderer, err := hashicorpraft.NewHashicorpRaftOrderer(
		nodeID,
		raftAddr,
		10*time.Second,
		"data/hermes/hashicor-raft/"+nodeID,
		2,
		10*time.Second,
		*listenJoinAddr,
		*joinAddr,
	)
	// hashicoprRaftOrderer.raft.leaderID = "localhost"
	// hashicoprRaftOrderer.raft.leaderID = "localhost:8000"
	if err != nil {
		log.Fatal(err.Error())
	}

	hermes := proxy.NewHermesProxy(httpCommunicator, hashicoprRaftOrderer)

	err = hermes.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}

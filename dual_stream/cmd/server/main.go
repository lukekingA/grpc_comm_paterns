package main

import (
	"net"

	"github.com/sirupsen/logrus"
	ts "gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/grpc_servers/dual_stream/internal/talkingserver"
	talkingservice "gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/grpc_servers/dual_stream/pkg/talkingservice/gen/go"
	"google.golang.org/grpc"
)

func main() {
	log := logrus.New()

	svrListen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to generate server listener: %v", err)
	}

	svr := grpc.NewServer()

	tss := ts.TalkingService{
		Log: log,
	}
	talkingservice.RegisterTalkingServiceServer(svr, &tss)

	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- svr.Serve(svrListen)
	}()

	log.Info("Chat Server Running")
	for e := range serverErrors {
		if e != nil {
			log.Fatalf("server failure: %v", e)
		}
	}

}

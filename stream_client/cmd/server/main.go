package main

import (
	"net"
	"os"

	"github.com/sirupsen/logrus"
	hubservice "gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/grpc_servers/stream_client/internal/hubservice"
	hbsvc "gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/grpc_servers/stream_client/pkg/hubupdater/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := logrus.New()

	svr := grpc.NewServer()

	hs := hubservice.HubService{
		Log: log,
	}

	hbsvc.RegisterHubUpdaterServiceServer(svr, &hs)

	reflection.Register(svr)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Errorf("failed to initiate stream server listener: %v", err)
		os.Exit(2)
	}

	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- svr.Serve(lis)
	}()
	log.Info("Hub Update Server Running")
	for e := range serverErrors {
		if e != nil {
			log.Fatalf("server failure: %v", e)
		}
	}

}

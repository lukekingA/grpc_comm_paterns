package main

import (
	"net"
	"os"

	"github.com/sirupsen/logrus"
	namesvcsvr "gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/grpc_servers/stream_server/internal/nameservice"
	namesservice "gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/grpc_servers/stream_server/pkg/namesvc/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := logrus.New()

	svr := grpc.NewServer()

	ns := namesvcsvr.NameService{
		Log: log,
	}

	namesservice.RegisterNamesServiceServer(svr, &ns)

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
	log.Info("Name Server Running")
	for e := range serverErrors {
		if e != nil {
			log.Fatalf("server failure: %v", e)
		}
	}

}

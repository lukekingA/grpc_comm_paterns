package main

import (
	"net"

	"github.com/sirupsen/logrus"
	"gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/2021_10_13/grpc/internal/hellosvc"
	hellov1 "gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/2021_10_13/grpc/pkg/v1/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := logrus.New()
	server := grpc.NewServer()

	hs := &hellosvc.HelloService{}

	hellov1.RegisterHelloServiceServer(server, hs)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	reflection.Register(server)

	err = server.Serve(lis)
	if err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

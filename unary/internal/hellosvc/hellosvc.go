package hellosvc

import (
	"context"
	"fmt"

	hellov1 "gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/2021_10_13/grpc/pkg/v1/gen/go"
)

type HelloService struct {
	hellov1.UnimplementedHelloServiceServer
}

func (hs *HelloService) Hello(ctx context.Context, hello *hellov1.HelloRequest) (*hellov1.HelloResponse, error) {
	greeting := fmt.Sprintf("Hello %s", hello.Name)
	return &hellov1.HelloResponse{
		Greeting: greeting,
	}, nil
}

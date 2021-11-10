package main

import (
	"context"
	"math/rand"
	"time"

	faker "github.com/bxcodec/faker/v3"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	hbsvc "gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/grpc_servers/stream_client/pkg/hubupdater/gen/go"
	"google.golang.org/grpc"
)

func main() {
	log := logrus.New()

	client, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial names service: %v", err)
	}

	namesClient := hbsvc.NewHubUpdaterServiceClient(client)

	s := rand.NewSource(time.Now().UnixNano())
	rndm := rand.New(s)
	numEvents := rndm.Intn(100)

	hubUpdaterClient, err := namesClient.UpdateStatusEvents(context.Background())
	if err != nil {
		log.Errorf("failed to setup hub updater client: %v", err)
	}

	for i := 0; i <= numEvents; i++ {
		evt := hbsvc.UpdateStatusEventsRequest{
			Id:    uuid.NewV4().String(),
			XCord: float32(faker.Latitude()),
			YCord: float32(faker.Longitude()),
		}
		err := hubUpdaterClient.Send(&evt)
		if err != nil {
			log.Errorf("failed to send update %s: %v", evt.Id, err)
		}
	}

	resp, err := hubUpdaterClient.CloseAndRecv()
	if err != nil {
		log.Errorf("failed to close hub updater: %v", err)
	} else {
		log.Infof("Updated events %v", resp.UpdatedId)
	}

}

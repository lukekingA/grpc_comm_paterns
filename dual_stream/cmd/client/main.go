package main

import (
	"context"
	"fmt"
	"io"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/grpc_servers/dual_stream/pkg/models"
	talkingservice "gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/grpc_servers/dual_stream/pkg/talkingservice/gen/go"
	"google.golang.org/grpc"
)

func main() {
	log := logrus.New()

	clDial, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to initiate grpc dialer: %v", err)
	}

	chatClient := talkingservice.NewTalkingServiceClient(clDial)

	clientStream, err := chatClient.Chat(context.Background())
	if err != nil {
		log.Fatalf("failed to initiate chat client: %v", err)
	}

	users := []models.User{
		{
			Id:   "1",
			Name: "Frank Johnson",
		},
		{
			Id:   "2",
			Name: "Bill Williams",
		},
		{
			Id:   "3",
			Name: "Maybel Baker",
		},
	}

	responses := []models.Chat{}
	mesages := map[string]string{}

	stopRecieve := make(chan struct{})
	go func() {
		for {
			respMsg, err := clientStream.Recv()
			if err == io.EOF {
				close(stopRecieve)
				return
			} else if err != nil {
				log.Errorf("failed to recieve chat response (client): %v", err)
			}
			c := models.Chat{
				ChatId:   respMsg.ChatId,
				Response: respMsg.ServerReplied,
			}

			responses = append(responses, c)
		}
	}()

	for i := 0; i < len(users)*4; i++ {
		u := users[i%len(users)]
		nu := talkingservice.User{
			Id:   u.Id,
			Name: u.Name,
		}
		chatreq := talkingservice.ChatRequest{
			User:     &nu,
			UserSaid: fmt.Sprintf("Yo server. Greetings from %s. What you got?", u.Name),
			ChatId:   uuid.NewV4().String(),
		}

		mesages[chatreq.ChatId] = chatreq.UserSaid

		err := clientStream.Send(&chatreq)
		if err != nil {
			log.Errorf("failed to send chat %s (client): %v", chatreq.ChatId, err)
		}
	}

	err = clientStream.CloseSend()
	if err != nil {
		log.Errorf("failed to close send stream (client): %v", err)
	}
	<-stopRecieve

	for _, r := range responses {
		fmt.Printf("Chat Request: %s\nChat Response: %s\n", mesages[r.ChatId], r.Response)
	}
}

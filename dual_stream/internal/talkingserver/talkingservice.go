package talkingservice

import (
	"fmt"
	"io"

	faker "github.com/bxcodec/faker/v3"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/grpc_servers/dual_stream/pkg/models"
	talkingservice "gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/grpc_servers/dual_stream/pkg/talkingservice/gen/go"
)

type TalkingService struct {
	Log *logrus.Logger
	talkingservice.UnimplementedTalkingServiceServer
}

func (ts *TalkingService) Chat(talkStream talkingservice.TalkingService_ChatServer) error {
	userChats := map[string][]models.Chat{}
	for {
		reqChat, err := talkStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			ts.Log.Errorf("failed to recieve chat (server): %v", err)
			continue
		}

		randResp := faker.Sentence()

		c := models.Chat{
			ChatId:   reqChat.ChatId,
			Chat:     reqChat.UserSaid,
			Response: randResp,
		}

		respId := uuid.NewV4().String()
		err = talkStream.Send(&talkingservice.ChatResponse{
			User:          reqChat.User,
			ServerReplied: randResp,
			ChatId:        reqChat.ChatId,
			ResponseId:    respId,
		})
		if err != nil {
			ts.Log.Errorf("failed to respond to chat (server) for user %s on chat %s: %v", reqChat.User.Name, reqChat.ChatId, err)
		} else {
			userChats[reqChat.User.Name] = append(userChats[reqChat.User.Name], c)
		}
	}

	for u, cs := range userChats {
		fmt.Printf("User: %v\n", u)
		fmt.Println("Chats")
		for _, c := range cs {
			fmt.Printf("User chat: %s\n\tResponse: %s\n", c.Chat, c.Response)
		}
	}
	return nil
}

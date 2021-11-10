package namesvcsvr

import (
	"fmt"
	"math/rand"
	"time"

	faker "github.com/bxcodec/faker/v3"
	"github.com/sirupsen/logrus"
	namesservice "gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/grpc_servers/stream_server/pkg/namesvc/gen/go"
)

type NameService struct {
	Log *logrus.Logger
	namesservice.UnimplementedNamesServiceServer
}

func (ns *NameService) GetNames(req *namesservice.GetNamesRequest, stream namesservice.NamesService_GetNamesServer) error {

	s := rand.NewSource(time.Now().UnixNano())
	rndm := rand.New(s)
	numNames := rndm.Intn(100)

	ns.Log.Infoln("Sending Names")
	for i := 0; i <= numNames; i++ {
		n := fmt.Sprintf("%s %s", faker.FirstName(), faker.LastName())
		err := stream.Send(&namesservice.GetNamesResponse{
			Name: n,
		})

		if err != nil {
			ns.Log.Errorf("failed to send name %s", n)
		}
	}
	return nil
}

package hubservice

import (
	"io"

	"github.com/sirupsen/logrus"
	hbsvc "gitlab.gs.kount.com/kount/Luke/daily_notes/2021/learnanddevops/grpc_servers/stream_client/pkg/hubupdater/gen/go"
)

type HubService struct {
	Log *logrus.Logger
	hbsvc.UnimplementedHubUpdaterServiceServer
}

func (hs *HubService) UpdateStatusEvents(stream hbsvc.HubUpdaterService_UpdateStatusEventsServer) error {

	savedIds := []string{}
	for {
		event, err := stream.Recv()
		if err == io.EOF {
			resp := hbsvc.UpdateStatusEventsResponse{
				UpdatedId: savedIds,
			}
			return stream.SendAndClose(&resp)
		}
		if err != nil {
			hs.Log.Errorf("failed to receive event")
			continue
		}

		savedIds = append(savedIds, event.Id)
	}
}

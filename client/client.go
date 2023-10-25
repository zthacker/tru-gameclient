package client

import (
	"context"
	"gameclient/frontend"
	"gameclient/proto"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type GameClient struct {
	CurrentPlayer uuid.UUID
	Stream        proto.GameBackend_StreamClient
	View          *frontend.View
}

func NewGameClient() *GameClient {
	return &GameClient{}
}

func (gc *GameClient) Connect(grpcClient proto.GameBackendClient, playerID uuid.UUID, password string, playerName string) error {
	// Connect to server
	req := proto.ConnectRequest{
		Id:       playerID.String(),
		Name:     playerName,
		Password: password,
	}

	resp, err := grpcClient.Connect(context.Background(), &req)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(resp)

	// Initialize Stream with token
	header := metadata.New(map[string]string{"authorization": resp.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	stream, err := grpcClient.Stream(ctx)
	if err != nil {
		return err
	}

	gc.Stream = stream
	gc.CurrentPlayer = playerID

	return nil
}

func (gc *GameClient) Start() {
	//write a loop to handle changes like moves, etc

	go func() {
		for {
			resp, err := gc.Stream.Recv()
			if err != nil {
				return
			}

			//need a lock here?
			switch resp.GetAction().(type) {
			case *proto.Response_AddEntity:
			case *proto.Response_RemoveEntity:
			case *proto.Response_UpdateEntity:
			}
		}
	}()

	//for {
	//	logrus.Infof("Client: %s is Connected", gc.CurrentPlayer)
	//	time.Sleep(10 * time.Second)
	//}

}

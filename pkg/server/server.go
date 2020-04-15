package server

import (
	"context"
	"time"

	"github.com/phil0522/znote/pkg/notesmarket"
	"github.com/phil0522/znote/pkg/subcommands"
	pb "github.com/phil0522/znote/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// ZNoteServer serve the znote service.
type ZNoteServer struct {
	GrpcServer *grpc.Server
	serveCount int32
	market     *notesmarket.Market
}

// Initialize initialize the server
func (s *ZNoteServer) Initialize() {
	logrus.Debug("initialize znote manger.")
	s.market = notesmarket.GetNotesMarket()
}

func (s *ZNoteServer) ExecuteCommand(ctx context.Context, req *pb.ZNoteRequest) (*pb.ZNoteResponse, error) {
	logrus.WithField("counter", s.serveCount).WithField("req", req).Debug("get request")
	s.serveCount++

	for _, command := range subcommands.Commands {
		if command.Name == req.Command {
			resp := command.ResolveRequest(*req)
			return &resp, nil
		}
	}
	return &pb.ZNoteResponse{}, nil
}

// QuitServer quits the server.
func (s *ZNoteServer) QuitServer(ctx context.Context, in *pb.QuitServerRequest) (*pb.QuitServerResponse, error) {
	go func() {
		time.Sleep(time.Second * 1)
		s.GrpcServer.Stop()
	}()
	return &pb.QuitServerResponse{}, nil
}

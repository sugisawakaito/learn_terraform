package handler

import (
	"context"
	"log"
	"server/entity"
	pb "server/proto"
	"server/service"

	"github.com/bluele/go-timecop"

	"golang.org/x/xerrors"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type messengerRequest struct {
	pb.UnimplementedMessengerServer
	requests []*pb.MessageRequest
}

func RegisterMessengerServer(s *grpc.Server) {
	pb.RegisterMessengerServer(s, &messengerRequest{})
}

func (s *messengerRequest) GetMessages(_ *empty.Empty, stream pb.Messenger_GetMessagesServer) error {
	for _, r := range s.requests {
		if err := stream.Send(&pb.MessageResponse{Message: r.GetMessage()}); err != nil {
			return xerrors.Errorf("failed to send a message: %w", err)
		}
	}

	previousCount := len(s.requests)

	for {
		currentCount := len(s.requests)
		if previousCount < currentCount && currentCount > 0 {
			r := s.requests[currentCount-1]
			log.Printf("Sent: %v", r.GetMessage())
			if err := stream.Send(&pb.MessageResponse{Message: r.GetMessage()}); err != nil {
				return xerrors.Errorf("failed to send a message: %w", err)
			}
		}
		previousCount = currentCount
	}
}

func (s *messengerRequest) CreateMessage(ctx context.Context, r *pb.MessageRequest) (*pb.MessageResponse, error) {
	msg := r.GetMessage()
	now := timecop.Now()

	newR := &pb.MessageRequest{Message: msg + ": " + now.Format("2006-01-02 15:04:05")}
	s.requests = append(s.requests, newR)

	if err := service.NewMessengerService().CreateUser(ctx, entity.Db, msg); err != nil {
		return nil, xerrors.Errorf("failed to create a user: %w", err)
	}
	return &pb.MessageResponse{Message: r.GetMessage()}, nil
}

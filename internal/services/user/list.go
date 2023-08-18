package user

import (
	"context"
	"log"

	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (svc service) List(ctx context.Context, emp *emptypb.Empty) (*pb.ListResponse, error) {
	users, err := svc.appCtx.UserRepository.List()
	if err != nil {
		log.Printf("Error unable to retrieve user list: %s\n", err.Error())
		return &pb.ListResponse{}, nil
	}

	return &pb.ListResponse{
		Ok:    true,
		Users: users.PB(),
	}, nil
}

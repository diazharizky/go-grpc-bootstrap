package user

import (
	"context"
	"log"

	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (svc service) List(ctx context.Context, emp *emptypb.Empty) (*pb.ListResponse, error) {
	mUsers, err := svc.appCtx.UserRepository.List()
	if err != nil {
		log.Printf("Error unable to retrieve user list: %s\n", err.Error())
		return &pb.ListResponse{}, nil
	}

	users := make([]*pb.User, len(mUsers))
	for i, mu := range mUsers {
		users[i].Username = mu.Username
		users[i].FullName = mu.FullName
		users[i].Email = mu.Email
	}

	return &pb.ListResponse{
		Ok:    true,
		Users: users,
	}, nil
}

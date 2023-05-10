package users

import (
	"context"
	"log"

	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (srv server) List(ctx context.Context, emp *emptypb.Empty) (*pb.ListResponse, error) {
	users, err := srv.appCtx.UserRepository.List()
	if err != nil {
		log.Printf("Error unable to retrieve user list: %s\n", err.Error())
		return &pb.ListResponse{}, nil
	}

	userList := pb.UserList{
		Users: users,
	}

	return &pb.ListResponse{
		Ok:    true,
		Users: &userList,
	}, nil
}

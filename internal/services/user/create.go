package user

import (
	"context"

	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

func (svc service) Create(ctx context.Context, newUser *pb.CreateParams) (*pb.CreateResponse, error) {
	user := &pb.User{
		Username: newUser.Username,
		FullName: newUser.FullName,
		Email:    newUser.Email,
	}

	if err := svc.appCtx.UserRepository.Create(user); err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		Ok:   true,
		User: user,
	}, nil
}

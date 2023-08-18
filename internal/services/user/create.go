package user

import (
	"context"

	"github.com/diazharizky/go-grpc-bootstrap/models"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

func (svc service) Create(ctx context.Context, newUser *pb.CreateParams) (*pb.CreateResponse, error) {
	user := &pb.User{
		Username: newUser.Username,
		FullName: newUser.FullName,
		Email:    newUser.Email,
	}

	mUser := &models.User{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}

	if err := svc.appCtx.UserRepository.Create(mUser); err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		Ok:   true,
		User: user,
	}, nil
}

package user

import (
	"context"

	"github.com/diazharizky/go-grpc-bootstrap/models"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

func (svc service) Create(ctx context.Context, newUser *pb.CreateParams) (*pb.CreateResponse, error) {
	mUser := &models.User{
		Username: newUser.Username,
		FullName: newUser.FullName,
		Email:    newUser.Email,
	}

	if err := svc.appCtx.UserRepository.Create(mUser); err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		Ok:   true,
		User: mUser.PB(),
	}, nil
}

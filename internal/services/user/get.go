package user

import (
	"context"

	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

func (svc service) Get(ctx context.Context, getParam *pb.GetUsernameParam) (*pb.GetResponse, error) {
	mUser, err := svc.appCtx.UserRepository.Get(getParam.Username)
	if err != nil {
		return nil, err
	}

	user := &pb.User{
		Username: mUser.Username,
		FullName: mUser.FullName,
		Email:    mUser.Email,
	}

	return &pb.GetResponse{
		Ok:   true,
		User: user,
	}, nil
}

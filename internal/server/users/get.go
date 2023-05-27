package users

import (
	"context"

	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

func (srv server) Get(ctx context.Context, getParam *pb.GetUsernameParam) (*pb.GetResponse, error) {
	user, err := srv.appCtx.UserRepository.Get(getParam.Username)
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{
		Ok:   true,
		User: user,
	}, nil
}

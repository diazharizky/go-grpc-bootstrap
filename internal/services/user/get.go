package user

import (
	"context"

	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

func (svc service) Get(ctx context.Context, getParam *pb.GetUsernameParam) (*pb.GetResponse, error) {
	user, err := svc.appCtx.UserRepository.Get(getParam.Username)
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{
		Ok:   true,
		User: user.PB(),
	}, nil
}

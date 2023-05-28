package user

import (
	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

type service struct {
	pb.UnimplementedUserServiceServer
	appCtx *app.Context
}

func NewService(appCtx *app.Context) *service {
	return &service{
		appCtx: appCtx,
	}
}

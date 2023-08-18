package user

import (
	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

type service struct {
	pb.UnimplementedUserServiceServer

	appCtx *app.Ctx
}

func NewService(appCtx *app.Ctx) *service {
	return &service{
		appCtx: appCtx,
	}
}

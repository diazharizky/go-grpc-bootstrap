package users

import (
	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

type server struct {
	pb.UnimplementedUserServiceServer
	appCtx *app.Context
}

func NewServer(appCtx *app.Context) *server {
	return &server{
		appCtx: appCtx,
	}
}

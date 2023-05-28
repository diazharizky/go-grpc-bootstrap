package users

import (
	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

type serviceServer struct {
	pb.UnimplementedUserServiceServer
	appCtx *app.Context
}

func NewServiceServer(appCtx *app.Context) *serviceServer {
	return &serviceServer{
		appCtx: appCtx,
	}
}

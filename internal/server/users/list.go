package users

import (
	"context"

	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server) List(ctx context.Context, emp *emptypb.Empty) (*pb.ListResponse, error) {
	return &pb.ListResponse{}, nil
}

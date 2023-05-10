package app

import (
	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

type IUserRepository interface {
	List() ([]*pb.User, error)
	Get() ([]*pb.User, error)
	Create(params pb.User) error
	Update(params pb.User) error
	Delete(username string) error
}

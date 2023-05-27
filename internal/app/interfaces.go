package app

import (
	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

type IUserRepository interface {
	List() (users []*pb.User, err error)
	Get(username string) (user *pb.User, err error)
	Create(newUser *pb.User) error
	Update(user *pb.User) error
	Delete(username string) error
}

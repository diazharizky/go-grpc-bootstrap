package repositories

import (
	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

type userRepository struct{}

func NewUserRepository() userRepository {
	return userRepository{}
}

func (userRepository) List() ([]*pb.User, error) {
	userList := []*pb.User{
		{
			Username:  "aominedaiki",
			FirstName: "Aomine",
			LastName:  "Daiki",
			Email:     "daiki.aomine@knb.com",
			Age:       18,
		},
		{
			Username:  "kiseryota",
			FirstName: "Kise",
			LastName:  "Ryota",
			Email:     "ryota.kise@knb.com",
			Age:       19,
		},
	}

	return userList, nil
}

func (userRepository) Get() ([]*pb.User, error) {
	return []*pb.User{}, nil
}

func (userRepository) Create(params pb.User) error {
	return nil
}

func (userRepository) Update(params pb.User) error {
	return nil
}

func (userRepository) Delete(username string) error {
	return nil
}

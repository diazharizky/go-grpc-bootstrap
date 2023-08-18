package app

import "github.com/diazharizky/go-grpc-bootstrap/models"

type IUserRepository interface {
	List() (users models.Users, err error)
	Get(username string) (user *models.User, err error)
	Create(newUser *models.User) error
	Update(user *models.User) error
	Delete(username string) error
}

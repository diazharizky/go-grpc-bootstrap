package repositories

import (
	"github.com/diazharizky/go-grpc-bootstrap/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(client *mongo.Client) userRepository {
	return userRepository{
		db: client.Database(""),
	}
}

func (r userRepository) List() (users []models.User, err error) {
	return
}

func (r userRepository) Get(username string) (*models.User, error) {
	return nil, nil
}

func (r userRepository) Create(newUser *models.User) error {
	return nil
}

func (userRepository) Update(params *models.User) error {
	return nil
}

func (userRepository) Delete(username string) error {
	return nil
}

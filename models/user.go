package models

import (
	"time"

	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	FullName  string             `json:"fullName" bson:"fullName"`
	Email     string             `json:"email" bson:"email"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time         `json:"updatedAt" bson:"updatedAt,omitempty"`
	DeletedAt *time.Time         `json:"deletedAt" bson:"deletedAt,omitempty"`
}

type Users struct {
	Data []User
}

func (u User) PB() *pb.User {
	return &pb.User{
		Username: u.Username,
		FullName: u.FullName,
		Email:    u.Email,
	}
}

func (u Users) PB() (users []*pb.User) {
	users = make([]*pb.User, len(u.Data))
	for i, user := range u.Data {
		users[i] = user.PB()
	}

	return
}

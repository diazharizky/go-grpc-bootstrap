package repositories

import (
	"database/sql"
	"errors"

	"github.com/diazharizky/go-grpc-bootstrap/pb"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) userRepository {
	return userRepository{
		db: db,
	}
}

func (rep userRepository) List() (users []*pb.User, err error) {
	syn := "SELECT id, username, full_name, email FROM users"

	rows, err := rep.db.Query(syn)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	for rows.Next() {
		user := &pb.User{}

		if err = rows.Scan(&user.Id, &user.Username, &user.FullName, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return
}

func (userRepository) Get() (*pb.User, error) {
	return &pb.User{}, nil
}

func (userRepository) Create(params *pb.User) error {
	return nil
}

func (userRepository) Update(params *pb.User) error {
	return nil
}

func (userRepository) Delete(username string) error {
	return nil
}

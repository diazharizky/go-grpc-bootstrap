package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	syn := "SELECT id, username, full_name, email, created_at FROM users"

	rows, err := rep.db.Query(syn)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	for rows.Next() {
		user := &pb.User{}

		var createdAt time.Time
		if err = rows.Scan(&user.Id, &user.Username, &user.FullName, &user.Email, &createdAt); err != nil {
			return nil, err
		}

		user.CreatedAt = timestamppb.New(createdAt)

		users = append(users, user)
	}

	return
}

func (rep userRepository) Get(username string) (*pb.User, error) {
	user := &pb.User{}

	syn := "SELECT id, username, full_name, email FROM users WHERE username = $1"

	err := rep.db.QueryRow(syn, username).Scan(
		&user.Id,
		&user.Username,
		&user.FullName,
		&user.Email,
	)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return user, nil
}

func (rep userRepository) Create(newUser *pb.User) error {
	syn := "INSERT INTO users (username, full_name, email, created_at) VALUES ($1, $2, $3, $4) RETURNING id"

	now := time.Now()
	err := rep.db.
		QueryRow(
			syn,
			newUser.Username,
			newUser.FullName,
			newUser.Email,
			now,
		).
		Scan(&newUser.Id)

	if err != nil {
		return err
	}

	return nil
}

func (userRepository) Update(params *pb.User) error {
	return nil
}

func (userRepository) Delete(username string) error {
	return nil
}

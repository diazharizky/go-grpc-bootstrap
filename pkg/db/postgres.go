package db

import (
	"database/sql"
	"fmt"

	"github.com/diazharizky/go-grpc-bootstrap/config"
	_ "github.com/lib/pq" // Require for PostgreSQL driver
)

type postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SslMode  string
}

func init() {
	config.Global.SetDefault("db.host", "0.0.0.0")
	config.Global.SetDefault("db.port", "5432")
	config.Global.SetDefault("db.user", "root")
	config.Global.SetDefault("db.password", "")
	config.Global.SetDefault("db.database", "")
	config.Global.SetDefault("db.sslmode", "disable")
}

func NewPostgres() postgres {
	return postgres{
		Host:     config.Global.GetString("db.host"),
		Port:     config.Global.GetString("db.port"),
		User:     config.Global.GetString("db.user"),
		Password: config.Global.GetString("db.password"),
		DBName:   config.Global.GetString("db.database"),
		SslMode:  config.Global.GetString("db.sslmode"),
	}
}

func (pg postgres) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", pg.dsn())
	if err != nil {
		return nil, fmt.Errorf("error unable to connect to PostgreSQL DB: %v", err)
	}

	return db, nil
}

func (pg postgres) dsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pg.Host,
		pg.Port,
		pg.User,
		pg.Password,
		pg.DBName,
		pg.SslMode,
	)
}

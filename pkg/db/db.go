package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/diazharizky/go-grpc-bootstrap/config"
	"github.com/diazharizky/go-grpc-bootstrap/internal/enum"
)

type IDatabase interface {
	Connect() (*sql.DB, error)
}

var (
	dbType enum.DBType
)

func init() {
	config.Global.SetDefault("db.type", "postgres")

	dbType = enum.DBType(config.Global.GetString("db.type"))
}

func GetConnection() (conn *sql.DB) {
	db, err := getDB()
	if err != nil {
		log.Fatalf("Error unable to get database connection: %v", err)
	}

	conn, err = db.Connect()
	if err != nil {
		log.Fatalf("Error unable to get database connection: %v", err)
	}

	return
}

func getDB() (db IDatabase, err error) {
	if !dbType.Validate() {
		err = fmt.Errorf("error invalid database type: %s", dbType.String())
		return
	}

	switch dbType {
	case enum.DBTypePostgres:
		db = NewPostgres()
	default:
		db = NewPostgres()
	}

	return
}

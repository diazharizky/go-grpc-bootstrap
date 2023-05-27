package main

import (
	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/internal/repositories"
	"github.com/diazharizky/go-grpc-bootstrap/internal/server"
	"github.com/diazharizky/go-grpc-bootstrap/pkg/db"
)

func main() {
	appCtx := &app.Context{}

	db := db.GetConnection()

	appCtx.UserRepository = repositories.NewUserRepository(db)

	server.Serve(appCtx)
}

package main

import (
	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/internal/repositories"
	"github.com/diazharizky/go-grpc-bootstrap/internal/server"
)

func main() {
	appCtx := &app.Context{}

	appCtx.UserRepository = repositories.NewUserRepository()

	server.Serve(appCtx)
}

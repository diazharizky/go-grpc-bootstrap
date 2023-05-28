package main

import (
	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/internal/repositories"
	"github.com/diazharizky/go-grpc-bootstrap/internal/services"
	"github.com/diazharizky/go-grpc-bootstrap/pkg/db"
)

func main() {
	appCtx := &app.Context{}

	db := db.GetConnection()

	appCtx.UserRepository = repositories.NewUserRepository(db)

	services.Serve(appCtx)
}

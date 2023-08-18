package main

import (
	"log"

	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/internal/repositories"
	"github.com/diazharizky/go-grpc-bootstrap/internal/services"
	"github.com/diazharizky/go-grpc-bootstrap/pkg/mongodb"
)

func main() {
	appCtx := &app.Ctx{}

	client, err := mongodb.GetClient()
	if err != nil {
		log.Fatalf("error unable to get db client: %v", err)
	}

	appCtx.UserRepository = repositories.NewUserRepository(client)

	services.Serve(appCtx)
}

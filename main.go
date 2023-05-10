package main

import (
	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/internal/server"
)

func main() {
	appCtx := &app.Context{}

	server.Serve(appCtx)
}

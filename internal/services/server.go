package services

import (
	"fmt"
	"log"
	"net"

	"github.com/diazharizky/go-grpc-bootstrap/config"
	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/internal/services/user"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var srv *grpc.Server

func init() {
	config.Global.SetDefault("server.host", "0.0.0.0")
	config.Global.SetDefault("server.port", "5000")
}

func Serve(appCtx *app.Ctx) {
	srv = grpc.NewServer()

	userServer := user.NewService(appCtx)

	pb.RegisterUserServiceServer(srv, userServer)

	reflection.Register(srv)

	addr := fmt.Sprintf("%s:%s", config.Global.GetString("server.host"), config.Global.GetString("server.port"))

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error unable to listen net address: %v\n", err)
	}

	log.Printf("Server is running on %v\n", listener.Addr())

	if err = srv.Serve(listener); err != nil {
		log.Fatalf("Error unable to run the server: %v\n", err)
	}
}

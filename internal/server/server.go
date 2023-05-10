package server

import (
	"fmt"
	"log"
	"net"

	"github.com/diazharizky/go-grpc-bootstrap/config"
	"github.com/diazharizky/go-grpc-bootstrap/internal/server/users"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var srv *grpc.Server

func init() {
	srv = grpc.NewServer()

	userServer := users.NewServer()

	pb.RegisterUserServiceServer(srv, userServer)

	reflection.Register(srv)

	addr := fmt.Sprintf("%s:%s", config.Global.GetString("server.host"), config.Global.GetString("server.port"))

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error unable to listen net address: %v\n", err)
	}

	log.Printf("Server running at %v\n", listener.Addr())

	if err = srv.Serve(listener); err != nil {
		log.Fatalf("Error unable to run the server: %v\n", err)
	}
}

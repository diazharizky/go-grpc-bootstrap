package users

import (
	"context"
	"log"
	"net"

	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type server struct {
	pb.UnimplementedUserServiceServer
	appCtx *app.Context
}

func NewServer(appCtx *app.Context) *server {
	return &server{
		appCtx: appCtx,
	}
}

func testServer(ctx context.Context, appCtx *app.Context) (pb.UserServiceClient, func()) {
	srv := grpc.NewServer()

	pb.RegisterUserServiceServer(srv, NewServer(appCtx))

	buffer := 101024 * 1024
	listener := bufconn.Listen(buffer)

	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Printf("Error unable to serve server: %v\n", err)
		}
	}()

	conn, err := grpc.DialContext(
		ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Printf("Error unable to connect to server: %v\n", err)
	}

	client := pb.NewUserServiceClient(conn)

	closer := func() {
		if err := listener.Close(); err != nil {
			log.Printf("Error unable to close listener: %v\n", err)
		}

		srv.Stop()
	}

	return client, closer
}

package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dkeng/cogo/proto/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type gServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *gServer) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Println(in.Name)
	fmt.Println(in.GetName())
	fmt.Println(in.String())
	return &helloworld.HelloReply{
		Message: "Server Message " + in.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	var gServer gServer
	helloworld.RegisterGreeterServer(s, &gServer)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

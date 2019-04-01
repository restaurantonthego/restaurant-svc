package main

import (
	"context"
	"log"
	"net"

	pb "github.com/brandon2255p/restaurant/restaurant-svc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Hello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.Response{Msg: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

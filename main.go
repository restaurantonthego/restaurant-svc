package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/looplab/eventhorizon/commandhandler/bus"
	eventbus "github.com/looplab/eventhorizon/eventbus/local"
	eventstore "github.com/looplab/eventhorizon/eventstore/mongodb"
	"github.com/restaurantonthego/restaurant-svc/domain"
	pb "github.com/restaurantonthego/restaurant-svc/proto"
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
	mongoURL := os.Getenv("MONGO_HOST")

	es, err := eventstore.NewEventStore(mongoURL, "restaurant")
	if err != nil {
		panic(fmt.Sprintf("Unable to establish eventstore - %s", err))
	}

	eb := eventbus.NewEventBus(nil)
	// go func() {
	// 	for err := range eb.Errors() {
	// 		log.Printf("EB: %s", err.Error())
	// 	}
	// }
	ch := bus.NewCommandHandler()
	domain.Setup(es, eb, ch)
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

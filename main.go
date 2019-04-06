package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/google/uuid"
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
type server struct {
	commandBus *bus.CommandHandler
}

// SayHello implements helloworld.GreeterServer
func (s *server) Hello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", in.Name)

	id, _ := uuid.Parse("25eb3db9-5833-11e9-b3ed-ac35ee1f3de8")
	err := s.commandBus.HandleCommand(ctx, &domain.ChangeRestaurantNameCommand{
		ID:   id,
		Name: in.Name,
	})
	if err != nil {
		return &pb.Response{Msg: "Hello we had an error"}, nil
	}
	return &pb.Response{Msg: "Hello " + in.Name}, nil

}

func main() {
	mongoURL := os.Getenv("MONGO_HOST")
	if mongoURL == "" {
		mongoURL = "localhost"
	}

	es, err := eventstore.NewEventStore(mongoURL, "restaurant")
	if err != nil {
		panic(fmt.Sprintf("Unable to establish eventstore - %s", err.Error()))
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
	pb.RegisterGreeterServer(s, &server{commandBus: ch})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

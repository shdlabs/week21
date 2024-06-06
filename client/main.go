package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/shdlabs/week21/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	id := flag.Int("id", 0, "User ID")
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("asus:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := service.NewQueryUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &service.UserRequest{Id: int32(*id)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Print(r.String())
}

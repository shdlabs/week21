package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	h "github.com/shdlabs/week21/helpers"
	"github.com/shdlabs/week21/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	id := flag.Int("id", 0, "User ID")
	host := flag.String("host", "localhost", "Server address")
	port := flag.String("port", "50051", "Server port")

	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.NewClient(
		strings.Join([]string{*host, *port}, ":"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := service.NewQueryUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	defer h.LogDuration(time.Now())

	r, err := c.GetUser(ctx, &service.UserRequest{Id: int32(*id)})
	if err != nil {
		log.Fatalf(h.Ko("could not greet: %v"), err)
	}

	log.Print(h.Ok(r.String()))
}

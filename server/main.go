package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	h "github.com/shdlabs/week21/helpers"
	"github.com/shdlabs/week21/service"
	"google.golang.org/grpc"
)

type server struct {
	service.UnimplementedQueryUserServer
	db service.DbMock
}

func (s *server) GetUser(ctx context.Context, in *service.UserRequest) (*service.UserReply, error) {
	log.Printf("Received: %v", in.GetId())

	defer h.LogDuration(time.Now())

	u, ok := s.db[in.GetId()]
	if !ok {
		return nil, errors.New("no such user ID")
	}

	return &service.UserReply{
		Id: u.ID, Fname: u.Fname,
		City: u.City, Phone: u.Phone,
		Height: u.Height, Married: u.Married,
	}, nil
}

func main() {
	db := mockTheData()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	service.RegisterQueryUserServer(s, &server{db: db})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func mockTheData() service.DbMock {
	db := service.NewDBMock()
	db.NewUsers(
		service.NewUser("John", "NY", "123456789", 1.75, true),
		service.NewUser("Steve", "LA", "123456789", 1.75, true),
		service.NewUser("Bill", "LA", "123456789", 1.75, true),
		service.NewUser("Joe", "LA", "123456789", 1.75, true),
	)
	return db
}

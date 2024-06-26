package main

import (
	"context"
	"fmt"
	"net"
	"time"

	h "github.com/shdlabs/week21/helpers"

	"github.com/charmbracelet/log"
	"github.com/shdlabs/week21/service"
	"google.golang.org/grpc"
)

type server struct {
	service.UnimplementedQueryUserServer
	db service.DbMock
}

func (s *server) GetUser(ctx context.Context, in *service.UserRequest) (*service.UserReply, error) {
	log.Info("Received: %v", in.GetId())

	defer h.DurationLog(time.Now(), "GetUser")

	u := s.db.FindUser(in.GetId())

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
		service.NewUser("John", "NY", "123456789", 1.75, false),
		service.NewUser("Anne", "CF", "123456789", 1.95, true),
		service.NewUser("Bill", "TA", "123456789", 1.85, false),
		service.NewUser("Alex", "LA", "123456789", 1.75, true),
		service.NewUser("Mark", "NY", "123456789", 1.75, false),
		service.NewUser("Mike", "NY", "123456789", 1.75, true),
		service.NewUser("Lisa", "NY", "123456789", 1.75, false),
		service.NewUser("Kate", "NY", "123456789", 1.75, true),
	)

	return db
}

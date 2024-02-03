package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "proto/gen/go/user/v1"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50055, "The server port")
)

type server struct {
	pb.UnimplementedUserServiceServer
}

// stub password login
func (s *server) PasswordLogin(ctx context.Context, in *pb.PasswordLoginRequest) (*pb.PasswordLoginResponse, error) {
	log.Printf("Received: %v, %v", in.GetEmail(), in.GetPassword())
	return &pb.PasswordLoginResponse{Email: "sample@example.com", UserId: "1", Admin: false}, nil
}

func StartServer() *grpc.Server {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return s
}

func main() {
	StartServer()
}

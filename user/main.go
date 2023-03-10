package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "proto/user"
)

var (
	port = flag.Int("port", 50055, "The server port")
)

type server struct {
	pb.UnimplementedUserServer
}

// stub password login
func (s *server) PasswordLogin(ctx context.Context, in *pb.PasswordLoginRequest) (*pb.LoginReply, error) {
	log.Printf("Received: %v, %v", in.GetEmail(), in.GetPassword())
	return &pb.LoginReply{Email: "sample@example.com", UserId: "1", Token: "abcdef", Expiry: "2025-01-01"}, nil
}

func StartServer() *grpc.Server {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return s
}

func main() {
	StartServer()
}

//package server
//
//import (
//	"flag"
//	"fmt"
//	"log"
//	"net"
//)
//
//import (
//	"google.golang.org/grpc"
//	pb "proto/user/user"
//)
//
//var (
//	port = flag.Int("port", 50051, "The server port")
//)
//
//type server struct {
//	pb.UnimplementedUserServiceServer
//}
//
//// stub password login
//func (s *server) PasswordLogin(ctx context.Context, in *pb.PasswordLoginRequest) (*pb.LoginReply, error) {
//	log.Printf("Received: %v", in.GetName())
//	return &pb.LoginReply{Message: "Hello " + in.GetName()}, nil
//}
//
//func StartServer() *grpc.Server {
//	flag.Parse()
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer()
//	pb.RegisterUserServiceServer(s, &server{})
//	log.Printf("server listening at %v", lis.Addr())
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//	return s
//}

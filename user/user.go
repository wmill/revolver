package main

import (
	"context"
	"fmt"
	"log"
	pb "proto/gen/go/user/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) UserCreate(ctx context.Context, in *pb.UserCreateRequest) (*pb.UserCreateResponse, error) {
	out := new(pb.UserCreateResponse)
	hashedPassword, err := HashPassword(in.GetPassword())
	if err != nil {
		log.Printf("Hashing error: %s", err)
	}
	org_id := 1
	var id int;
	err = GetDbConn().QueryRow("INSERT INTO users (email, password_hash, admin, org_id) VALUES ($1, $2, $3, $4) RETURNING id", in.GetEmail(), hashedPassword, in.GetAdmin(), org_id).Scan(&id)
	if err != nil {
		log.Printf("Insert error: %s", err)
		return nil, err
	}
	out.Admin = false
	out.Email = in.GetEmail()
	out.UserId = fmt.Sprint(id)
	return out, nil
}

func (s *server) PasswordLogin(ctx context.Context, in *pb.PasswordLoginRequest) (*pb.PasswordLoginResponse, error) {
	log.Printf("Received: %v, %v", in.GetEmail(), in.GetPassword())
	var id int
	var email string
	var admin bool
	var hashedPassword string
	// hashedPassword, err := HashPassword(in.GetPassword())
	// if err != nil {
	// 	log.Printf("Hashing error: %s", err)
	// 	return nil, err
	// }
	err := GetDbConn().QueryRow("SELECT id, email, admin, password_hash FROM users WHERE email = $1", in.GetEmail()).Scan(&id, &email, &admin, &hashedPassword)
	if err != nil {
		log.Printf("Query error: %s", err)
		return nil, err
	}
	if !CheckPasswordHash(in.GetPassword(), hashedPassword) {
		log.Printf("Invalid password")
		return &pb.PasswordLoginResponse{}, status.Errorf(codes.PermissionDenied,"Invalid password")
	}
	return &pb.PasswordLoginResponse{Email: email, UserId: fmt.Sprint(id), Admin: admin}, nil
	// return &pb.PasswordLoginResponse{Email: "sample@example.com", UserId: "1", Admin: false}, nil
}
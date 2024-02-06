package main

import (
	"context"
	"fmt"
	"log"
	pb "proto/gen/go/user/v1"
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
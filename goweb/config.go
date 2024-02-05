package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	UserServGrpc  string
	JwtSecret     string
	JwtExpiration time.Duration
	Port					string
}

var globalConfig Config

func LoadConfig() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	userServGrpc := os.Getenv("USER_SERV_GRPC")
	if userServGrpc == "" {
		log.Fatal("USER_SERV_GRPC is not set")
	}
	jwtExpiration, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOURS"))
	if err != nil {
		log.Fatal("JWT_EXPIRATION_HOURS is not set")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set")
	}
	globalConfig = Config{
		UserServGrpc:  userServGrpc,
		JwtSecret:     secret,
		JwtExpiration: time.Duration(jwtExpiration),
		Port: port,
	}
}

func GetConfig() Config {
	return globalConfig
}

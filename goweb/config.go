package main

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	UserServGrpc  string `env:"USER_SERV_GRPC, required"`
	JwtSecret     string `env:"JWT_SECRET, required"`
	JwtExpiration time.Duration `env:"JWT_EXPIRATION_HOURS, required"`
	Port					string `env:"PORT, required"`
}

var globalConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
    log.Fatal(err)
  }
	ctx := context.Background()
	err = envconfig.Process(ctx, &globalConfig)
	if err != nil {
    log.Fatal(err)
  }
}

func GetConfig() Config {
	return globalConfig
}

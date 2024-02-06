package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)


type Config struct {
	UserDbHost string `env:"USER_DB_HOST, required"`
	UserDbPort int    `env:"USER_DB_PORT, required"`
	UserDbUser string `env:"USER_DB_USER, required"`
	UserDbPass string `env:"USER_DB_PASSWORD, required"`
	UserDbName string `env:"USER_DB_NAME, required"`
	DbSSLMode  string `env:"DB_SSL_MODE, required"`
	Port       int 	  `env:"PORT, required"`
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



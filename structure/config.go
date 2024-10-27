package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)


type Config struct {
	StructureDbHost string `env:"STRUCTURE_DB_HOST, required"`
	StructureDbPort int    `env:"STRUCTURE_DB_PORT, required"`
	StructureDbUser string `env:"STRUCTURE_DB_USER, required"`
	StructureDbPass string `env:"STRUCTURE_DB_PASSWORD, required"`
	StructureDbName string `env:"STRUCTURE_DB_NAME, required"`
	DbSSLMode  string `env:"DB_SSL_MODE, required"`
	Port       int 	  `env:"PORT, required"`
}

var globalConfig Config

func LoadConfig() {

	err := godotenv.Load()
	if err != nil {
    log.Print(err)
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



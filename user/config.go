package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


type Config struct {
	UserDbHost string
	UserDbPort int
	UserDbUser string
	UserDbPass string
	UserDbName string
	DbSSLMode  string
	Port			 int
}

var globalConfig Config

func LoadConfig() {

	err := godotenv.Load()

	if err != nil {
    log.Fatal("Error loading .env file")
  }
	
	userDbHost := os.Getenv("USER_DB_HOST")
	if userDbHost == "" {
		log.Fatal("USER_DB_HOST is not set")
	}
	userDbPort, err := strconv.Atoi(os.Getenv("USER_DB_PORT"))
	if err != nil {
		log.Fatal("USER_DB_PORT is not set")
	}
	userDbUser := os.Getenv("USER_DB_USER")
	if userDbUser == "" {
		log.Fatal("USER_DB_USER is not set")
	}
	userDbPass := os.Getenv("USER_DB_PASSWORD")
	if userDbPass == "" {
		log.Fatal("USER_DB_PASSWORD is not set")
	}
	userDbName := os.Getenv("USER_DB_NAME")
	if userDbName == "" {
		log.Fatal("USER_DB_NAME is not set")
	}
	dbSSLMode := os.Getenv("DB_SSL_MODE")
	if userDbName == "" {
		log.Fatal("DB_SSL_MODE is not set")
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("PORT is not set")
	}
	globalConfig = Config{
		UserDbHost: userDbHost,
		UserDbPort: userDbPort,
		UserDbUser: userDbUser,
		UserDbPass: userDbPass,
		UserDbName: userDbName,
		DbSSLMode: dbSSLMode,
		Port: port,
	}
}

func GetConfig() Config {
	return globalConfig
}

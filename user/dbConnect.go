package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var dbConn *sql.DB

func ConnectToDb() {
	config := GetConfig()
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=%s",

	
	config.UserDbHost, config.UserDbPort, config.UserDbUser, config.UserDbPass, config.UserDbName, config.DbSSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("UserDb connection error",err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("UserDb ping error",err)
	}
	fmt.Println("UserDb connected")
	dbConn = db
}

func GetDbConn() *sql.DB {
	return dbConn
}

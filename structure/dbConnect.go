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

	
	config.StructureDbHost, config.StructureDbPort, config.StructureDbUser, config.StructureDbPass, config.StructureDbName, config.DbSSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("StructureDb connection error",err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("StructureDb ping error",err)
	}
	fmt.Println("StructureDb connected")
	dbConn = db
}

func GetDbConn() *sql.DB {
	return dbConn
}

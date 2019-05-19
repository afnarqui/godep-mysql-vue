package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

//connection should conect with bd
func getConnection() *sql.DB {
	dsn := "postgress://golang:golang@127.0.0.1:5432/gocrud?sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}
	
	return db
}

package main

import (
	"./src/system/app"
	DB "learning-golang/api.example.com/src/system/db"
	"flag"
	"os"
	"github.com/joho/godotenv"
	
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "Assigning the port that thel server should listen on.")

	flag.Parse()
	err := godotenv.Load("config.ini")

	if err != nil {
		panic(err)
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}

func main() {
	db, err := DB.Connect()
	if err != nil {
		panic(err)
	}
	s := app.NewServer()

	s.Init(port, db)
	s.Start()
}


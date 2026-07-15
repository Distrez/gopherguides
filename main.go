package main

import (
	"gopherguides/connection"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := connection.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Connected successfully")
}

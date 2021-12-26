package main

import (
	"log"
	"os"

	"TodoPS/db"
	"TodoPS/handler"
	"TodoPS/repository"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}
	database, err := db.NewDB(db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(database.GetDB())
	handlers := handler.NewHandler(repo)
	if err = handlers.Run(); err != nil {
		log.Fatal(err)
	}
}

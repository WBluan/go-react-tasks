package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func GetMongoUri() string {
	return os.Getenv("MONGODB_URI")
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	return port
}

package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	loadDotEnv()
	return os.Getenv("MONGOURI")
}

func GetDatabaseName() string {
	loadDotEnv()
	return os.Getenv("DATABASE_NAME")
}

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

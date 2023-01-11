package infrastructure

import (
	"log"
    "fmt"

	"github.com/joho/godotenv"

)

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
     fmt.Println("env loaded succesfully")
}

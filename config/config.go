package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	//return os.Getenv(key)
}

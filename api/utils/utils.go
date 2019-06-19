package utils

import "github.com/joho/godotenv"

// LoadEnv will load environment variables
func LoadEnv() {
	godotenv.Load()
}

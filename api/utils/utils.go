package utils

import (
	env "github.com/joho/godotenv"
)

// LoadEnv should load .env file
func LoadEnv() {
	env.Load()
}

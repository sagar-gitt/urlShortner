package config

import (
	"github.com/joho/godotenv"
	"log"
)

const (
	e1      = "Error loading .env file"
	envPath = "../.env"
)

func LoadEnv() {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal(e1, err)
	}
}

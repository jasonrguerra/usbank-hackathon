package main

import (
	"github.com/joho/godotenv"
	"errors"
	"os"
)

type Auth struct {
	ApiKey     string
	Secret     string
	CustomerID string
}

func loadEnv() (Auth, error) {
	err := godotenv.Load("private.env")
	if err != nil {
		return Auth{}, errors.New("No env file found")
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return Auth{}, errors.New("No api key found")
	}
	secrets := os.Getenv("SECRET")
	if secrets == "" {
		return Auth{}, errors.New("No secret found")
	}
	customerID := os.Getenv("customerID")
	if customerID == "" {
		return Auth{}, errors.New("No customer id found")
	}
	authentication := Auth{apiKey, secrets, customerID}
	return authentication, nil

}


package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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

type Category struct {
	EatingOut   float64
	Groceries   float64
	Streaming   float64
	Insurance   float64
	Charity     float64
	ShoppingRec float64
	Car         float64
	Misc        float64
}

// take a struct* and add num to it
func catagorize(subcategory string, amount float64, data Category) Category {
	num, _ := strconv.Atoi(subcategory)
	switch num {
	case 5499, 5812, 5814:
		data.EatingOut += amount
	case 5411:
		data.Groceries += amount
	case 4899:
		data.Streaming += amount
	case 6300:
		data.Insurance += amount
	case 8398:
		data.Charity += amount
	case 5651, 5611, 7278, 7999:
		data.ShoppingRec += amount
	case 5541, 7538:
		data.Car += amount
	case 5732, 5719:
		data.Misc += amount
	default:
		fmt.Println("unaccounted code ", num)
		data.Misc += amount
	}
	return data
}

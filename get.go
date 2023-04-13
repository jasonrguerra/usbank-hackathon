package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"encoding/json"
)

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	AccountID        string   `json:"accountID"`                   
	PaymentAccountID string  `json:"paymentAccountID,omitempty"`  
	AccOuntType      string  `json:"acc\nountType,omitempty"`     
	Description      string   `json:"description"`                 
	Status           string   `json:"status"`                      
	AvailableCredit  float64 `json:"available\nCredit,omitempty"` 
	CurrentBalance   float64 `json:"currentBalance,omitempty"`    
	PaymentACcountID string  `json:"paymentA\nccountID,omitempty"`
	AccountType      string  `json:"accountType,omitempty"`       
	AvailableBalance float64 `json:"availableBalance,omitempty"`  
	Nickname         string  `json:"nickname,omitempty"`          
	AccoUntType      string  `json:"acco\nuntType,omitempty"`     
}

func main() {
	err := godotenv.Load("private.env")
	if err != nil {
		log.Fatal("private.env file not found")
	}

	customerID := os.Getenv("customerID")
	if customerID == "" {
		log.Fatal("no customer id found")
	}
	requestURL := "https://alpha-api.usbank.com/innovation/bank-node/customer-accounts/v1/customer/" + customerID + "/accounts"
	apiKey := os.Getenv("API_KEY")
	secret := os.Getenv("SECRET")
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Fatal("Unable to make request")
	}
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
	}
	req.SetBasicAuth(apiKey, secret)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed do")
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var data Accounts
	if err := json.Unmarshal(b, &data); err != nil {
		fmt.Println("Failed to unmarshal JSON")
	}
	fmt.Println(data.Accounts[1].CurrentBalance)




//	fmt.Println(string(b))

}

package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"fmt"
)
const baseUrl = "https://alpha-api.usbank.com/innovation/bank-node/customer-accounts/v1/"

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
//	Nickname         string  `json:"nickname,omitempty"`
	AccountID        string  `json:"accountID"`
//	PaymentAccountID string  `json:"paymentAccountID"`
//	AccountType      string  `json:"accountType"`
//	Description      string  `json:"description"`
//	Status           string  `json:"status"`
//	AvailableBalance float64 `json:"availableBalance,omitempty"`
//	LineOfCredit     float64 `json:"lineOfCredit,omitempty"`
//	AvailableCredit  float64 `json:"availableCredit,omitempty"`
//	CurrentBalance   float64 `json:"currentBalance,omitempty"`
}

type Balances struct {
	Balances []Balance `json:"accounts"`
}

type Balance struct {
	Description      string  `json:"description"`
	AvailableBalance   float64 `json:"availableBalance,omitempty"`
	CurrentBalance   float64 `json:"currentBalance,omitempty"`
}

type Transactions struct {
	Transactions []Transaction 
}

type Transaction struct {
	AccountID              string `json:"accountID"`
	Amount                 int64  `json:"amount"`
	CardNumber             string `json:"cardNumber"`
	Category               string `json:"category"`
	Channel                string `json:"channel"`
	CheckNumber            string `json:"checkNumber"`
	DebitCreditMemo        string `json:"debitCreditMemo"`
	Description            string `json:"description"`
	ForeignAmount          int64  `json:"foreignAmount"`
	ForeignCurrency        string `json:"foreignCurrency"`
	MaskedCardNumber       string `json:"maskedCardNumber"`
	Memo                   string `json:"memo"`
	Payee                  string `json:"payee"`
	PostedTimestamp        string `json:"postedTimestamp"`
	Reference              string `json:"reference"`
	ReferenceTransactionID string `json:"referenceTransactionID"`
	Status                 string `json:"status"`
	Subcategory            string `json:"subcategory"`
	TransactionID          string `json:"transactionID"`
	TransactionTimestamp   string `json:"transactionTimestamp"`
	TransactionType        string `json:"transactionType"`
}



func getAccountsID(auth Auth) ([]int, error) {

	requestURL := baseUrl + "customer/" + auth.CustomerID + "/accounts"
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, errors.New("Unable to make request")
	}
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
	}
	req.SetBasicAuth(auth.ApiKey, auth.Secret)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Failed do")
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var data Accounts
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, errors.New("Failed to unmarshal JSON")
	}
	var accountSlice []int
	for i := 0; i < len(data.Accounts); i++ {
		num, _ := strconv.Atoi(data.Accounts[i].AccountID)
		accountSlice = append(accountSlice, num)
	}
	return accountSlice, nil
}

func getTransactions(accountID int, auth Auth) (string,error) {
	id := strconv.Itoa(accountID)
	requestURL := baseUrl + "/account/" + id + "/trans/" + "SHORT"
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return "", errors.New("Unable to make request")
	}
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
	}
	req.SetBasicAuth(auth.ApiKey, auth.Secret)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", errors.New("Failed do")
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var data Transactions
	if err := json.Unmarshal(b, &data); err != nil {
		return "", errors.New("Failed to unmarshal JSON")
	}
	fmt.Println(data)

	return string(b),nil
}

func getBalances(auth Auth) (Balances, error) {

	requestURL := baseUrl + "customer/" + auth.CustomerID + "/accounts"
	req, err := http.NewRequest("GET", requestURL, nil)
	var data Balances
	if err != nil {
		return data, errors.New("Unable to make request")
	}
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
	}
	req.SetBasicAuth(auth.ApiKey, auth.Secret)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return data, errors.New("Failed do")
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(b, &data); err != nil {
		return data, errors.New("Failed to unmarshal JSON")
	}


	return data, nil


}

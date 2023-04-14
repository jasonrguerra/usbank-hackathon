package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
)

const baseUrl = "https://alpha-api.usbank.com/innovation/bank-node/customer-accounts/v1/"

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	//	Nickname         string  `json:"nickname,omitempty"`
	//AccountID string `json:"accountID"`
	PaymentAccountID string `json:"paymentAccountID"`
	// AccountType      string  `json:"accountType"`
	// Description      string  `json:"description"`
	// Status           string  `json:"status"`
	// AvailableBalance float64 `json:"availableBalance,omitempty"`
	// LineOfCredit     float64 `json:"lineOfCredit,omitempty"`
	// AvailableCredit  float64 `json:"availableCredit,omitempty"`
	// CurrentBalance   float64 `json:"currentBalance,omitempty"`
}

type Balances struct {
	Balances []Balance `json:"accounts"`
}

type Balance struct {
	Description      string  `json:"description"`
	AvailableBalance float64 `json:"availableBalance,omitempty"`
	CurrentBalance   float64 `json:"currentBalance,omitempty"`
}

type Transactions struct {
	Transactions []Transaction
}

type Transaction struct {
	AccountID              string  `json:"accountID"`
	TransactionID          string  `json:"transactionID"`
	CardNumber             string  `json:"cardNumber,omitempty"`
	MaskedCardNumber       string  `json:"maskedCardNumber,omitempty"`
	ReferenceTransactionID string  `json:"referenceTransactionID,omitempty"`
	PostedTimestamp        string  `json:"postedTimestamp"`
	TransactionTimestamp   string  `json:"transactionTimestamp"`
	Channel                string  `json:"channel"`
	Description            string  `json:"description"`
	Memo                   string  `json:"memo,omitempty"`
	DebitCreditMemo        string  `json:"debitCreditMemo"`
	Category               string  `json:"category,omitempty"`
	Subcategory            string  `json:"subcategory,omitempty"`
	Reference              string  `json:"reference,omitempty"`
	Status                 string  `json:"status"`
	Amount                 float64 `json:"amount"`
	Payee                  string  `json:"payee,omitempty"`
	CheckNumber            string  `json:"checkNumber,omitempty"`
	ForeignAmount          float64 `json:"foreignAmount,omitempty"`
	ForeignCurrency        string  `json:"foreignCurrency,omityempty"`
	TransactionType        string  `json:"transactionType"`
}

func getAccountsIDsFromBank(auth Auth) ([]int, error) {
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
		num, _ := strconv.Atoi(data.Accounts[i].PaymentAccountID)
		accountSlice = append(accountSlice, num)
	}
	return accountSlice, nil
}

func getTransactionsFromBank(accountID int, auth Auth) (Transactions, error) {
	id := strconv.Itoa(accountID)
	requestURL := baseUrl + "/account/" + id + "/trans/" + "PUR"
	req, err := http.NewRequest("GET", requestURL, nil)
	var data Transactions
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
		os.Exit(1)
		return data, errors.New("Failed to unmarshal JSON")

	}

	return data, nil
}

func getBalancesfromBank(auth Auth) (Balances, error) {

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

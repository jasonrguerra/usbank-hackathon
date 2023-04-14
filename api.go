package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPing(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}

func getBalances(c *gin.Context) {
	auth, err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}
	balances, err := getBalancesfromBank(auth)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Unable to get balance from bank"})
		return
	}
	c.IndentedJSON(http.StatusOK, balances)
}

func getTransactions(c *gin.Context) {
	auth, err := loadEnv()
	id := c.Param("id")

	if err != nil {
		log.Fatal(err)
	}

	accountID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id must be a integer"})
	}
	transactions, err := getTransactionsFromBank(accountID, auth)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found on server"})
		return
	}
	c.IndentedJSON(http.StatusOK, transactions)
}

func getAccountIDs(c *gin.Context) {
	auth, err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}
	accountIDs, err := getAccountsIDsFromBank(auth)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Unable to get balance from bank"})
		return
	}
	c.IndentedJSON(http.StatusOK, accountIDs)
}

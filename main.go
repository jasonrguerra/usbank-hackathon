package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", getPing)
	r.GET("/balances", getBalances)
	r.GET("/accountid", getAccountIDs)
	r.GET("/transactions/:id", getTransactions)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

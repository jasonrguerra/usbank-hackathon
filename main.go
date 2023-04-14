package main

import (
//	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func main() {
//	auth, err := loadEnv()
//	if err != nil {
//		log.Fatal(err)
//	}
	//balances,err := getBalances(auth)
	//if err != nil {
	//	log.Fatal(err)
	//}
//	fmt.Println(balances)



	r := gin.Default()
	r.GET("/ping",getPing)
	r.GET("balances",GetBalances)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


func getPing(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
	"message": "pong",
	})

}

func GetBalances(c *gin.Context) { 
	auth, err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}
	balances, _ := getBalances(auth)
    c.IndentedJSON(http.StatusOK, balances)
}

package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/*
	const data = {
	    type: 'doughnut',
	    data: {
	    labels: ['Red', 'Blue', 'Yellow'],
	    datasets: [{
	        label: '# of Votes',
	        data: [12, 19, 3, 5, 2, 3],
	        borderWidth: 1
	    }]
	 },

}
*/

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", getPing)
	r.GET("/balances", getBalances)
	r.GET("/accountid", getAccountIDs)
	r.GET("/transactions/:id", getTransactions)
	r.GET("/categories", getCategory)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

package main 

import (
		"net/http"
		"github.com/gin-gonic/gin"
		
)

type kpw struct {
	Date	string	`json:"date"`
	Key		bool	`json:"key"`
	Phone	bool	`json:"phone"`	
	Wallet	bool	`json:"wallet"`
	Notes	string	`json:"notes"`
}

var kpws = []kpw{
	{Date: "01/01/2023", Key: false, Phone: false, Wallet: false, Notes: "notes"},
	{Date: "01/02/2023", Key: false, Phone: false, Wallet: false, Notes: "notes"},
	{Date: "01/03/2023", Key: false, Phone: false, Wallet: false, Notes: "notes"},
}	

func getKpws(c *gin.Context){
	c.IndentedJSON(http.StatusOK, kpws)
}

func main() {
	router:= gin.Default()
	router.Get("/kpws", getKpws)
	router.Run("localhost:9090")
}
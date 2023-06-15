package main

import (
	"golang/model"
	"golang/module/transaction"

	"github.com/gin-gonic/gin"
)

func main() {
	model.ConnectDatabase()
	r := gin.Default()
	r.GET("/get-transactions", transaction.GetAllTransactions)
	r.GET("/get-transactions/:id/:end", transaction.GetAllTransactionsByParam)
	r.GET("/getTransactions/:status", transaction.GetTransactionByStatus)
	r.GET("/getTransactions2/:date", transaction.FindByDate)
	r.Run()
}

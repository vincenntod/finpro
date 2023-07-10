package main

import (
	"golang/model"
	"golang/module/transaction"

	"github.com/gin-gonic/gin"
)

func main() {
	model.ConnectDatabase()

	r := gin.Default()

	//Clean-code

	transactionHandler := transaction.DefaultRequestHandler(model.DB)
	r.GET("/get-transactions/", transactionHandler.GetAllTransaction)
	r.GET("/get-transactions-status/:status/", transactionHandler.GetAllTransactionByStatus)
	r.GET("/get-transactions-date/:start/:end/", transactionHandler.GetAllTransactionByDate)
	r.GET("/get-transactions-statusdate/:status/:start/:end/", transactionHandler.GetAllTransactionByStatusDate)
	r.GET("/get-transactions-limit/:id", transactionHandler.GetAllLimit)

	r.Run()
}

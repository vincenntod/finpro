package main

import (
	"golang/auth"
	"golang/model"
	"golang/module/account"
	"golang/module/exportcsv"
	"golang/module/transaction"

	"github.com/gin-gonic/gin"
)

func main() {
	model.ConnectDatabase()
	r := gin.Default()
	export := exportcsv.DefaultRequestHandler(model.DB)
	accountHandler := account.DefaultRequestHandler(model.DB)

	Admin := r.Group("/api", auth.MiddlewareAdmin)
	{
		//table transaction
		Admin.GET("/get-transactions", transaction.GetAllTransactions)
		Admin.GET("/get-transactions/:id/:end", transaction.GetAllTransactionsByParam)
		Admin.GET("/getTransactions/:status", transaction.GetTransactionByStatus)
		Admin.GET("/getTransactions2/:date", transaction.FindByDate)

		//export CSV
		Admin.GET("/export-transaction", export.ExportCSVHandler)

		//account
		Admin.GET("/data-user", accountHandler.GetDataUser)
		Admin.GET("/data-user/:id", accountHandler.GetDataUserById)
		Admin.PUT("/data-user/:id", accountHandler.EditDataUser)
		Admin.POST("/create-user", accountHandler.CreateAccount)
		Admin.DELETE("/data-user/:id", accountHandler.DeleteDataUser)
		Admin.GET("/logout", accountHandler.Logout)
	}
	r.POST("/login", accountHandler.Login)
	r.Run(":8081")

}

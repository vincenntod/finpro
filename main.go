package main

import (
	"golang/auth"
	"golang/model"
	"golang/module/account"
	"golang/module/exportcsv"
	"golang/module/transactions"

	"github.com/gin-gonic/gin"
)

func main() {
	model.ConnectDatabase()
	r := gin.Default()
	r.Use(auth.CORSMiddleware())
	export := exportcsv.DefaultRequestHandler(model.DB)
	accountHandler := account.DefaultRequestHandler(model.DB)

	Admin := r.Group("/api", auth.MiddlewareAdmin)
	{
		//account
		Admin.GET("/data-user", accountHandler.GetDataUser)
		Admin.GET("/data-user/:id", accountHandler.GetDataUserById)
		Admin.PUT("/data-user/:id", accountHandler.EditDataUser)
		Admin.DELETE("/data-user/:id", accountHandler.DeleteDataUser)
		Admin.GET("/logout", accountHandler.Logout)

		//export csv file
		Admin.GET("/export-transaction", export.ExportCSVHandler)

		//transaction table
		Admin.GET("/get-transactions", transactions.GetAllTransactions)
		Admin.GET("/get-transactions-limit/:id", transactions.GetAllTransactionsLimit)
		Admin.GET("/getTransactions/:status", transactions.GetAllTransactionByStatus)
		Admin.GET("/getTransactionsDate/:start/:end", transactions.GetAllTransactionByDate)
		Admin.GET("/getTransactionsStatusDate/:status/:start/:end", transactions.GetAllTransactionByStatusDate)

	}
	r.GET("/export-transaction", export.ExportCSVHandler)
	r.POST("/create-user", accountHandler.CreateAccount)
	r.POST("/login", accountHandler.Login)
	r.Run(":8081")

}

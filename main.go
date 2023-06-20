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
		Admin.GET("/export-transaction/status/:status", export.ExportCSVHandlerStatusfilter)
		Admin.GET("/export-transaction/range-date/:start_date/:end_date", export.ExportCSVHandlerRangeDateFilter)
		Admin.GET("/export-transaction/status-range-date/:status/:start_date/:end_date", export.ExportCSVHandlerStatusAndRangeDateFilter)

		//transaction table
		Admin.GET("/get-transactions", transactions.GetAllTransactions)
		Admin.GET("/get-transactions-limit/:id", transactions.GetAllTransactionsLimit)
		Admin.GET("/getTransactions/:status", transactions.GetAllTransactionByStatus)
		Admin.GET("/getTransactionsDate/:start/:end", transactions.GetAllTransactionByDate)
		Admin.GET("/get-transaction-status-date/:id", transactions.GetTransactionByStatusDate)

	}

	r.POST("/login", accountHandler.Login)
	r.Run(":8081")
}

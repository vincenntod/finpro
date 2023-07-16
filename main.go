package main

import (
	"golang/auth"
	"golang/model"
	"golang/module/account"
	"golang/module/exporttransaction"
	"golang/module/transaction"

	"github.com/gin-gonic/gin"
)

func main() {
	model.ConnectDatabase()
	r := gin.Default()
	r.Use(auth.CORSMiddleware())
	export := exporttransaction.DefaultRequestHandler(model.DB)
	accountHandler := account.DefaultRequestHandler(model.DB)
	transactionHandler := transaction.DefaultRequestHandler(model.DB)

	Admin := r.Group("/api", auth.MiddlewareAdmin)
	{
		//account
		Admin.GET("/data-user", accountHandler.GetDataUser)
		Admin.GET("/data-user/:id", accountHandler.GetDataUserById)
		Admin.PUT("/data-user/:id", accountHandler.EditDataUser)
		Admin.DELETE("/data-user/:id", accountHandler.DeleteDataUser)

		//export csv file
		Admin.GET("/export-transaction/csv", export.ExportTransactionHandler)

		//transaction table
		Admin.GET("/get-transactions/", transactionHandler.GetAllTransaction)
		Admin.GET("/get-transaction-status/:status/", transactionHandler.GetAllTransactionByStatus)
		Admin.GET("/get-TransactionDate/:start/:end/", transactionHandler.GetAllTransactionByDate)
		Admin.GET("/get-TransactionStatusDate/:status/:start/:end/", transactionHandler.GetAllTransactionByStatusDate)
		Admin.GET("/get-transactions-limit/:id", transactionHandler.GetAllLimit)

	}

	r.POST("/create-user", accountHandler.CreateAccount)
	r.POST("/send-email-registration/:email", accountHandler.SendEmailRegister)

	//forgot password
	r.POST("/send-email/:email", accountHandler.SendEmail)
	r.POST("/compare-verification-code", accountHandler.CompareVerificationCode)
	r.PUT("/edit-password/", accountHandler.EditPassword)

	r.POST("/login", accountHandler.Login)
	r.Run()
}

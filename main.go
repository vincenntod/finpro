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
		Admin.GET("/get-transaction-by-status", transactionHandler.GetAllTransactionByRequestLimit)


	}
	//create new user
	r.POST("/create-user", accountHandler.CreateAccount)
	r.POST("/send-email-registration/:email", accountHandler.SendEmailRegistration)

	//forgot password
	r.POST("/send-email-forgot-password/:email", accountHandler.SendEmailForgotPassword)
	r.POST("/compare-verification-code", accountHandler.CompareVerificationCode)
	r.PUT("/edit-password/", accountHandler.EditPassword)

	r.POST("/login", accountHandler.Login)
	r.Run()
}

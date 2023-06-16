package main

import (
	"golang/auth"
	"golang/model"
	"golang/module/account"
	"golang/module/exportcsv"

	"github.com/gin-gonic/gin"
)

func main() {
	model.ConnectDatabase()
	export := exportcsv.DefaultRequestHandler(model.DB)
	accountHandler := account.DefaultRequestHandler(model.DB)
	r := gin.Default()

	Admin := r.Group("/api", auth.MiddlewareAdmin)
	{
		Admin.GET("/export-transaction", export.ExportCSVHandler)
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

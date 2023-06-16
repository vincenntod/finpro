package main

import (
	"golang/auth"
	"golang/model"
	"golang/module/account"

	"github.com/gin-gonic/gin"
)

func main() {
	model.ConnectDatabase()
	r := gin.Default()
	accountHandler := account.DefaultRequestHandler(model.DB)

	Admin := r.Group("/api", auth.MiddlewareAdmin)

	{
		Admin.GET("/data-user", accountHandler.GetDataUser)
		Admin.GET("/data-user/:id", accountHandler.GetDataUserById)
		Admin.PUT("/data-user/:id", accountHandler.EditDataUser)
		Admin.POST("/create-user", accountHandler.CreateAccount)
		Admin.DELETE("/data-user/:id", accountHandler.DeleteDataUser)
		Admin.GET("/logout", accountHandler.Logout)
	}
	r.POST("/login", accountHandler.Login)
	r.Run(":8080")
}

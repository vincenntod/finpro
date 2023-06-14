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
	Admin := r.Group("/api", auth.MiddlewareAdmin)
	{
		Admin.GET("/data-user", account.GetDataUser)
		Admin.GET("/data-user/:id", account.GetDataUserById)
		Admin.PUT("/data-user/:id", account.EditDataUser)
		Admin.DELETE("/data-user/:id", account.DeleteDataUser)
	}
	r.POST("/create-user", account.CreateAccount)
	r.POST("/login", account.Login)
	r.Run(":8080")
}

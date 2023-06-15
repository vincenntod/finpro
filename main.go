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
	usersHandler := account.DefaultRequestHandler(model.DB)

	Admin := r.Group("/api", auth.MiddlewareAdmin)

	{
		Admin.GET("/data-user", usersHandler.GetDataUser)
		Admin.GET("/data-user/:id", usersHandler.GetDataUserById)
		Admin.PUT("/data-user/:id", usersHandler.EditDataUser)
		Admin.POST("/create-user", usersHandler.CreateAccount)
		Admin.DELETE("/data-user/:id", usersHandler.DeleteDataUser)
		Admin.GET("/logout", usersHandler.Logout)
	}
	r.POST("/login", usersHandler.Login)
	r.Run(":8080")
}

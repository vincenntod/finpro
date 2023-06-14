package main

import (
	"golang/model"

	"github.com/gin-gonic/gin"
)

func main() {
	model.ConnectDatabase()
	r := gin.Default()
	// Admin := r.Group("/api", auth.MiddlewareAdmin)
	{
	}

	r.Run(":8080")
}

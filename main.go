package main

import (
	"golang/model"
	"golang/module/exportcsv"

	"github.com/gin-gonic/gin"
)

func main() {
	model.ConnectDatabase()
	export := exportcsv.DefaultRequestHandler(model.DB)
	r := gin.Default()
	// Admin := r.Group("/api", auth.MiddlewareAdmin)
	{
	}
	r.GET("/export-transaction", export.ExportCSVHandler)

	r.Run(":8080")
}

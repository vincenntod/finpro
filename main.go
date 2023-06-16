package main

import (
	"golang/model"
	"golang/module/transactionNotClean"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func initDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=12345 dbname=dbceria port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
}

func main() {
	model.ConnectDatabase()

	//Clean-code
	//db, err := initDB()
	//if err != nil {
	//	log.Fatal(err)
	//}

	r := gin.Default()

	//Clean-code
	//transactionHandler := transaction.DefaultRequestHandler(db)
	//r.GET("/get-transaction-status/:status", transactionHandler.GetTransactionByStatus)

	r.GET("/get-transactions", transactionNotClean.GetAllTransactions)

	r.GET("/get-transactions-limit/:id", transactionNotClean.GetAllTransactionsLimit)

	r.GET("/getTransactions/:status", transactionNotClean.GetAllTransactionByStatus)
	
	r.GET("/get-transaction-date/:id", transactionNotClean.GetTransactionByDate)

	r.GET("/get-transaction-status-date/:id", transactionNotClean.GetTransactionByStatusDate)
	r.Run()
}

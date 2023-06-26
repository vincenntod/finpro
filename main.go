package main

import (
	"golang/model"
	"golang/module/transaction"

	"github.com/gin-gonic/gin"
)

/*
func initDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=12345 dbname=dbceria port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
}*/

func main() {
	model.ConnectDatabase()

	//Clean-code
	// db, err := initDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	r := gin.Default()

	//Clean-code

	transactionHandler := transaction.DefaultRequestHandler(model.DB)
	r.GET("/get-transactions", transactionHandler.GetAllTransaction)
	r.GET("/get-transaction-status/:status", transactionHandler.GetAllTransactionByStatus)
	r.GET("/getTransactionDate/:start/:end", transactionHandler.GetAllTransactionByDate)
	r.GET("/getTransactionStatusDate/:status/:start/:end", transactionHandler.GetAllTransactionByStatusDate)

	//Gak dipake
	/*
		r.GET("/get-transactions-date/:id", transactionHandler.GetTransactionByDate)
		r.GET("/get-transaction-status-date/:id", transactionHandler.GetTransactionByStatusAndDate)
	*/
	//Not Clean-code
	// r.GET("/getTransactionDate/:start/:end", transactionNotClean.GetAllTransactionByDate)
	/*
		r.GET("/getTransactionsStatusDate/:status/:start/:end", transactionNotClean.GetAllTransactionByStatusDate)

		r.GET("/get-transactions", transactionNotClean.GetAllTransactions)

		r.GET("/get-transactions-limit/:id", transactionNotClean.GetAllTransactionsLimit)

		r.GET("/get-status", transactionNotClean.GetTransactionByQueryStatus)

		r.GET("/get-transaction-date/:id", transactionNotClean.GetTransactionByDate)

		r.GET("/get-transaction-status-date/:id", transactionNotClean.GetTransactionByStatusDate)
	*/
	r.Run()
}

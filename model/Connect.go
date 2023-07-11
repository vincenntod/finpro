package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=123456 dbname=dbceria port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		fmt.Printf("Error")
		return
	}
	DB = db
}

// func ConnectDatabase() {
// 	dsn := "root:12345@tcp(localhost:3306)/dbceria?parseTime=true"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
// 		SingularTable: true,
// 	}})
// 	if err != nil {
// 		fmt.Printf("Error")
// 		return
// 	}
// 	DB = db
// }

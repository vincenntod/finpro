package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "host=ep-muddy-boat-134191.us-east-2.aws.neon.tech user=IchsanR password=v7jfgMB6hEPW dbname=neondb port=5432 sslmode=verify-full TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		fmt.Printf("Error")
		return
	}
	DB = db
}

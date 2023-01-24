package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	connectionString := "user:password@tcp(localhost:8080)/database?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())

	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

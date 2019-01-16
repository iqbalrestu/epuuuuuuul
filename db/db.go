package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

//Init ...
func Init() {
	var err error
	db, err = gorm.Open("mysql", "root@tcp(localhost:3306)/epuuuuuuul?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)

}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

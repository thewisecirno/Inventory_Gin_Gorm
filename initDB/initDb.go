package initDB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	Db *gorm.DB
)

func InitDB() *gorm.DB {
	dsn := "root:qq31415926535--@tcp(localhost:3306)/listofitems?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err2 := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{})
	if err2 != nil {
		log.Fatal(err2)
	}
	return DB
}

package initDB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	Db  *gorm.DB
	err error
)

func InitDB() {
	dsn := "root:qq31415926535--@tcp(localhost:3306)/listofitems?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

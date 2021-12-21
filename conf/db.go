package conf

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// var (
	// 	username = "root"
	// 	password = ""
	// 	hostname = "127.0.0.1"
	// 	dbname   = "mygram_golang"
	// 	port     = "3306"
	// 	dsn      = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, port, dbname)
	// )

	dsn := "root:@tcp(127.0.0.1:3306)/mygram-golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(dsn)
		panic(err.Error())
	}

	return db
}

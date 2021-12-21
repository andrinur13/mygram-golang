package conf

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var (
		username = "root"
		password = "root"
		hostname = "127.0.0.1"
		dbname   = "mygram_golang"
		dsn      = fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}

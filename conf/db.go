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

	userName := "andri"
	password := "ServerAzure123@"
	hostIP := "20.124.39.116"
	portServer := "3306"
	dbName := "mygram-golang"

	// read db
	dsnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, hostIP, portServer, dbName)

	// dsn := "root:@tcp(127.0.0.1:3306)/mygram-golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsnString), &gorm.Config{})
	if err != nil {
		fmt.Println(dsnString)
		panic(err.Error())
	}

	return db
}

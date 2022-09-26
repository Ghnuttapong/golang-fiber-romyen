package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/mysql"	
	"gorm.io/gorm"
)


var (
	Database *gorm.DB
)


func NewConnection() error {
	dsn := getURL()
	var err error
	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
			log.Panic("failed to connect database, have an error: ",err.Error())
	}
	return nil
}



func getURL() string {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Println("error on load db port from env: ", err.Error())
		port = 3306
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_HOST"),
		port,
		os.Getenv("DATABASE_NAME"))
}
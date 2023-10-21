package db

import (
	"backend-go/internal/pkg/config"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config.DataBase.UserName, config.Config.DataBase.Password, config.Config.DataBase.Host, config.Config.DataBase.DB_Name)), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql connect failed, err: ", err)
		os.Exit(1)
	}

	DB = db
}

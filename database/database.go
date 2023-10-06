package database

import (
	"fmt"
	"uts/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	
	dsn := "root@tcp(127.0.0.1:3306)/admin_management_user?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err:= gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	  fmt.Println("database tidak dapat terhubung")
	}
	
   fmt.Println("Database terkoneksi")
  
   DB = conn
   DB.AutoMigrate(models.User{}, models.Admin{})
}

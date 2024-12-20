package config

import (
	"fmt"
	"os"

	"github.com/afrizalsebastian/go-gin-gorm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db_user := os.Getenv("USER_DB")
	db_password := os.Getenv("PASSWORD_DB")
	db_url := os.Getenv("BASE_DB")
	db_port := os.Getenv("PORT_DB")
	db_name := os.Getenv("NAME_DB")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db_user, db_password, db_url, db_port, db_name)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	DB = db
}

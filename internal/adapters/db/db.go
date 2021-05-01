package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // mySql import
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ml-x-men/config"
	"ml-x-men/internal/application"
)

type DB struct {
	*gorm.DB
}

func New(c config.DBConfig) (application.Storage, error) {
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Name)
	db, err := gorm.Open(mysql.Open(conStr), &gorm.Config{})
	return &DB{
		db,
	}, err
}

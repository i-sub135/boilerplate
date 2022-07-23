package service

import (
	"fmt"
	"github.com/gocraft/dbr/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type DBConnect struct {
	Connection *dbr.Connection
	DSN        string
}

func GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}

func (po *DBConnect) Conn() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(po.DSN), &gorm.Config{})
	return db, err
}

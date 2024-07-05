package db

import (
	"fmt"
	"github.com/littlehole/paper-sharing/internal/db/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

type DatabaseConfig struct {
	username string
	password string
	host     string
	port     string
	database string
	location string
}

var (
	db   *gorm.DB
	once sync.Once
)

func NewDB() *gorm.DB {
	once.Do(initDB)
	return db
}

// 初始化数据库连接
func initDB() {
	var err error
	db, err = gorm.Open(mysql.Open(getDns()), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	db.AutoMigrate(&models.UserModel{})
	db.AutoMigrate(&models.LabModel{})
}

func getDns() string {
	dbconfig := DatabaseConfig{
		username: "root",
		password: "123456",
		host:     "127.0.0.1",
		port:     "3306",
		database: "paperSharing",
		location: "Local",
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		dbconfig.username, dbconfig.password, dbconfig.host, dbconfig.port, dbconfig.database, dbconfig.location)
}

package repository

import (
	"fmt"
	"github.com/brcodingdev/go-crud-users/internal/adapters/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	db *gorm.DB
)

// Connect connects to DB
func Connect() (*gorm.DB, error) {
	dbUserName := os.Getenv("DB_USERNAME")
	dbUserPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	metricsLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUserName, dbUserPassword, dbHost, dbPort, dbName,
	)
	fmt.Println(dsn)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: metricsLogger})

	if err != nil {
		return nil, err
	}
	db = d
	return db, nil
}

// MigrateDB migrates db following mapping Tables
func MigrateDB() error {
	return db.AutoMigrate(model.Tables...)
}

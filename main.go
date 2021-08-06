package main

import (
	"log"
	"os"
	"time"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DSN              = "root@tcp(127.0.0.1:3306)/user_db?charset=utf8mb4&parseTime=True&loc=Local"
	MIGRATION_FOLDER = "migrations"
)

func main() {
	log.Println("Server initialize")
	// init step
	time.Local = time.UTC
	// set log level
	// TODO

	// connect redis
	// TODO

	// connect DB
	log.Println("Database connection establish")
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: 7 * time.Second, // Slow SQL threshold
				LogLevel:      logger.Error,    // Log level
				Colorful:      true,            // Disable color
			},
		),
	})
	if err != nil {
		panic(err)
	}
	sql, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = sql.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Execute migration")
	// migrate
	goose.SetDialect("mysql")
	err = goose.Up(sql, MIGRATION_FOLDER)
	if err != nil {
		panic(err)
	}

	log.Println("Server running!")
}

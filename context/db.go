package context

import (
	"log"
	"os"
	"time"

	"github.com/allen012694/usersystem/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	gormDb, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{
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
		return nil, err
	}

	db = gormDb
	return gormDb, nil
}

func GetDB() *gorm.DB {
	return db
}

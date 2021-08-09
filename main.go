package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/allen012694/usersystem/common"
	"github.com/allen012694/usersystem/config"
	"github.com/pressly/goose/v3"
)

func main() {
	time.Local = time.UTC

	// set log level
	// TODO

	// init step
	log.Infoln("Server initialize")

	// connect redis
	_, err := common.InitRedis()
	if err != nil {
		panic(err)
	}
	log.Infoln("Redis connection established")

	// connect DB
	db, err := common.InitDB()
	if err != nil {
		panic(err)
	}
	log.Infoln("Database connection established")

	// migrate
	sql, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sql.Close()
	goose.SetDialect("mysql")
	err = goose.Up(sql, config.MIGRATION_FOLDER)
	if err != nil {
		panic(err)
	}

	server := &server{}
	server.Init(config.RUNNING_PORT)
	log.Infoln("Server running!!")
	if err := server.Serve(); err != nil {
		log.Errorln(err)
	}

	log.Infoln("Server stopeed!")
}

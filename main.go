package main

import (
	"log"
	"time"

	"github.com/allen012694/usersystem/config"
	"github.com/allen012694/usersystem/context"
	"github.com/pressly/goose/v3"
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
	db, err := context.InitDB()
	if err != nil {
		panic(err.Error())
	}
	log.Println("Database connection established")

	// migrate
	sql, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	defer sql.Close()
	goose.SetDialect("mysql")
	err = goose.Up(sql, config.MIGRATION_FOLDER)
	if err != nil {
		panic(err.Error())
	}

	server := &server{}
	server.Init(config.RUNNING_PORT)
	log.Println("Server running!!")
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

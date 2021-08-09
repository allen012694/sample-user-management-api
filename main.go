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
	_, err := context.InitRedis()
	if err != nil {
		panic(err)
	}
	log.Println("Redis connection established")

	// connect DB
	db, err := context.InitDB()
	if err != nil {
		panic(err)
	}
	log.Println("Database connection established")

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
	log.Println("Server running!!")
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Server stopeed!")
}

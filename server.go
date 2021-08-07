package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/allen012694/usersystem/routes"
	"github.com/gin-gonic/gin"
)

type server struct {
	port   string
	router *gin.Engine
}

func (server *server) Init(port string) {
	server.port = port

	server.router = gin.Default()
	server.router.GET("/ping", ping)
	server.router.GET("/", home)

	// API V1
	apiV1 := server.router.Group("/v1")
	apiV1.POST("/login", routes.Login)
	apiV1.PATCH("/users", routes.UpdateUser)
	apiV1.GET("/users/me", routes.GetCurrentUser)
	apiV1.POST("/assets/upload", routes.UploadAsset)
}

func (server *server) Serve() error {
	// run service
	return server.router.Run(fmt.Sprintf(":%v", server.port))
}

func home(ctx *gin.Context) {
	log.Println("Called '/'")
	ctx.JSON(http.StatusOK, gin.H{"ok": 1})
}

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

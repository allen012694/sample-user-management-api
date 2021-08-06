package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type server struct {
	Port   string
	Router *gin.Engine
}

func (server *server) Init(port string) {
	server.Port = port
	server.Router = gin.Default()

	server.Router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.Router.GET("/", home)
	server.Router.POST("login", login)
	server.Router.PATCH("/users", updateUser)
	server.Router.GET("/users/me", getCurrentUser)
	server.Router.POST("/assets/upload", uploadAsset)
}

func (server *server) Serve() error {
	// run service
	return server.Router.Run(fmt.Sprintf(":%v", server.Port))
}

// ROUTES
func home(ctx *gin.Context) {
	log.Println("Called '/'")
	ctx.JSON(http.StatusOK, gin.H{"ok": 1})
}
func login(ctx *gin.Context) {
	log.Println("Called '/login'")
	ctx.JSON(http.StatusOK, gin.H{"ok": 1})
}
func updateUser(ctx *gin.Context) {
	log.Println("Called '/users'")
	ctx.JSON(http.StatusOK, gin.H{"ok": 1})
}
func getCurrentUser(ctx *gin.Context) {
	log.Println("Called '/users/me'")
	ctx.JSON(http.StatusOK, gin.H{"ok": 1})
}
func uploadAsset(ctx *gin.Context) {
	log.Println("Called '/assets/upload'")
	ctx.JSON(http.StatusOK, gin.H{"ok": 1})
}

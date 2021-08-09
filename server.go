package main

import (
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/allen012694/usersystem/config"
	"github.com/allen012694/usersystem/models/user"
	"github.com/allen012694/usersystem/routes"
	"github.com/allen012694/usersystem/utils"
	"github.com/gin-gonic/gin"
)

type server struct {
	port   string
	router *gin.Engine
}

func (server *server) Init(port string) {
	server.port = port

	server.router = gin.New()
	// server.router.Use(gin.Logger())
	server.router.Use(gin.Recovery())
	server.router.Use(handleError)

	server.router.GET("/ping", ping)
	server.router.GET("/", home)

	// API V1
	apiV1 := server.router.Group("/v1")
	apiV1.POST("/login", routes.Login)

	apiV1.GET("/users/me", authRequire, routes.GetCurrentUser)
	apiV1.PATCH("/users/me", authRequire, routes.UpdateCurrentUser)
	apiV1.POST("/assets/upload", routes.UploadAsset)
}

func (server *server) Serve() error {
	// run service
	return server.router.Run(fmt.Sprintf(":%v", server.port))
}

func home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"ok": 1})
}

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func authRequire(ctx *gin.Context) {
	token, err := utils.ExtractJwtTokenFromHeaderString(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.AbortWithError(403, errors.New(config.ErrorLoginSessionInvalid))
		return
	}

	// Validate against SECRET of JWT
	tokeninzer := utils.NewJwtTokenizer()
	payload, err := tokeninzer.Validate(token)
	if err != nil {
		ctx.AbortWithError(403, errors.New(config.ErrorLoginSessionInvalid))
		return
	}

	// Check in redis store
	err = utils.CheckStoreSession(ctx.Request.Context(), token)
	if err != nil {
		ctx.AbortWithError(403, errors.New(config.ErrorLoginSessionInvalid))
		return
	}

	// Retrieve corresponding user
	user, err := user.GetUserById(payload.Id)
	if err != nil {
		ctx.AbortWithError(403, errors.New(config.ErrorUserNotExisted))
		return
	}

	ctx.Set(config.CONTEXT_CURRENT_USER, user)
	ctx.Next()
}

func handleError(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) > 0 {
		log.Errorln(ctx.Errors)
		ctx.JSON(ctx.Writer.Status(), ctx.Errors.JSON())
	}
}

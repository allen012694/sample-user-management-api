package routes

import (
	"log"
	"net/http"

	"github.com/allen012694/usersystem/controllers"
	"github.com/allen012694/usersystem/types"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	log.Println("Called '/login'")
	var request types.LoginRequest
	if err := ctx.ShouldBind(&request); err != nil {
		panic(err.Error())
	}

	response, err := controllers.Login(ctx.Request.Context(), &request)
	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(http.StatusOK, response)
}

func UpdateUser(ctx *gin.Context) {
	log.Println("Called '/users'")
	ctx.JSON(http.StatusOK, gin.H{"ok": 1})
}

func GetCurrentUser(ctx *gin.Context) {
	log.Println("Called '/users/me'")
	ctx.JSON(http.StatusOK, gin.H{"ok": 1})
}

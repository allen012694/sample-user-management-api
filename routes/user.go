package routes

import (
	"net/http"

	"github.com/allen012694/usersystem/config"
	"github.com/allen012694/usersystem/controllers"
	"github.com/allen012694/usersystem/models/user"
	"github.com/allen012694/usersystem/types"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var request types.LoginRequest
	if err := ctx.ShouldBind(&request); err != nil {
		panic(err)
	}

	response, err := controllers.Login(ctx.Request.Context(), &request)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, response)
}

func UpdateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"ok": 1})
}

func GetCurrentUser(ctx *gin.Context) {
	currentUser := ctx.MustGet(config.CURRENT_USER).(*user.User)
	ctx.JSON(http.StatusOK, currentUser)
}

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
		ctx.AbortWithError(500, err)
		return
	}

	response, err := controllers.Login(ctx.Request.Context(), &request)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func UpdateCurrentUser(ctx *gin.Context) {
	currentUser := ctx.MustGet(config.CONTEXT_CURRENT_USER).(*user.User)

	var request types.UpdateUserRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	request.UserId = currentUser.Id
	response, err := controllers.UpdateUser(ctx.Request.Context(), &request)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func GetCurrentUser(ctx *gin.Context) {
	currentUser := ctx.MustGet(config.CONTEXT_CURRENT_USER).(*user.User)
	ctx.JSON(http.StatusOK, currentUser)
}

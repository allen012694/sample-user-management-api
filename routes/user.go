package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/allen012694/usersystem/config"
	"github.com/allen012694/usersystem/controllers"
	"github.com/allen012694/usersystem/models/user"
	"github.com/allen012694/usersystem/types"
	"github.com/allen012694/usersystem/utils"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var request types.LoginRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := controllers.Login(ctx.Request.Context(), &request)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, response)

	go func() {
		jsonRequest, _ := json.Marshal(request)
		jsonResponse, _ := json.Marshal(response)
		log.Infoln(utils.ActivityLog{
			Subject:   "user",
			SubjectId: fmt.Sprint(response.UserId),
			Object:    "user",
			ObjectId:  fmt.Sprint(response.UserId),
			Action:    config.LOG_ACTION_LOGIN,
			Request:   string(jsonRequest),
			Response:  string(jsonResponse),
		})
	}()
}

func UpdateCurrentUser(ctx *gin.Context) {
	currentUser := ctx.MustGet(config.CONTEXT_CURRENT_USER).(*user.User)

	var request types.UpdateUserRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	request.UserId = currentUser.Id
	response, err := controllers.UpdateUser(ctx.Request.Context(), &request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, response)

	go func() {
		jsonRequest, _ := json.Marshal(request)
		jsonResponse, _ := json.Marshal(response)
		log.Infoln(utils.ActivityLog{
			Subject:   "user",
			SubjectId: fmt.Sprint(request.UserId),
			Object:    "user",
			ObjectId:  fmt.Sprint(request.UserId),
			Action:    config.LOG_ACTION_UPDATE,
			Request:   string(jsonRequest),
			Response:  string(jsonResponse),
		})
	}()
}

func GetCurrentUser(ctx *gin.Context) {
	currentUser := ctx.MustGet(config.CONTEXT_CURRENT_USER).(*user.User)
	ctx.JSON(http.StatusOK, currentUser)
}

func UploadCurrentUserProfilePicture(ctx *gin.Context) {
	currentUser := ctx.MustGet(config.CONTEXT_CURRENT_USER).(*user.User)

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Save the file (for now save file into local)
	filename := filepath.Base(file.Filename)
	container := filepath.Join("files", currentUser.Username)
	os.MkdirAll(container, 0700) // ensure folders existed
	destination := filepath.Join(container, filename)
	if err := ctx.SaveUploadedFile(file, destination); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// persist record into DB
	request := &types.AddProfilePictureRequest{
		UserId:     currentUser.Id,
		PictureUrl: destination,
	}
	response, err := controllers.AddProfilePicture(ctx.Request.Context(), request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, response)

	go func() {
		jsonRequest, _ := json.Marshal(request)
		jsonResponse, _ := json.Marshal(response)
		log.Infoln(utils.ActivityLog{
			Subject:   "user",
			SubjectId: fmt.Sprint(request.UserId),
			Object:    "profile_picture",
			ObjectId:  fmt.Sprint(response.Id),
			Action:    config.LOG_ACTION_CREATE,
			Request:   string(jsonRequest),
			Response:  string(jsonResponse),
		})
	}()
}

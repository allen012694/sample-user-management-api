package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadAsset(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"ok": 1})
}

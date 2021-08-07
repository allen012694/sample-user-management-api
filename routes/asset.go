package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadAsset(ctx *gin.Context) {
	log.Println("Called '/assets/upload'")
	ctx.JSON(http.StatusOK, gin.H{"ok": 1})
}

package http

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/files"
	"github.com/gin-gonic/gin"
)

func MapFileRoutes(fileGroups *gin.RouterGroup, fh files.Handlers) {
	fileGroups.POST("/upload", fh.UploadFile)
	fileGroups.GET("/:file_name", fh.GetFile)
	fileGroups.DELETE("/:file_name", fh.DeleteFile)
}

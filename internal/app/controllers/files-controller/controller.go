package filesController

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	UploadFile(c *gin.Context)
	GetFile(c *gin.Context)
	DeleteFile(c *gin.Context)
}

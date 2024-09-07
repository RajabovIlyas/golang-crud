package file

import "github.com/gin-gonic/gin"

type Handlers interface {
	UploadFile(g *gin.Context)
	GetFile(g *gin.Context)
	DeleteFile(g *gin.Context)
}

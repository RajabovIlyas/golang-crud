package filesController

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/services/files-service"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type FilesController struct {
	fs *filesService.FilesService
}

func NewFilesController(p *models.DBConfigParam) *FilesController {
	return &FilesController{fs: filesService.NewFilesService(p)}
}

func (fc *FilesController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}

	format := utils.GetFormatFile(file)
	path := utils.GetPathFileByFormat(format, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}

	filePath, err := fc.fs.UploadFile(c.Request.Context(), database.CreateFileParams{
		Format:   format,
		Path:     path,
		FileName: file.Filename,
		UserID:   uuid.NullUUID{},
		Size:     file.Size,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Message{
		Message: filePath,
	})
}

func (fc *FilesController) GetFile(c *gin.Context) {
	fileName := c.Param("file_name")
	file, err := fc.fs.FindFile(c.Request.Context(), fileName)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorModel{Error: "file not found"})
	}
	c.File(file.Path)
}

func (fc *FilesController) DeleteFile(c *gin.Context) {
	fileName := c.Param("file_name")
	err := fc.fs.DeleteFile(c.Request.Context(), fileName)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Message{Message: "File deleted"})

}

package http

import (
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/file"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type fileHandlers struct {
	cfg    *config.Config
	fileUC file.UseCase
}

func NewFileHandlers(cfg *config.Config, fileUC file.UseCase) file.Handlers {
	return &fileHandlers{cfg: cfg, fileUC: fileUC}
}

func (f fileHandlers) UploadFile(g *gin.Context) {
	newFile, err := g.FormFile("file")

	if err != nil {
		g.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}

	format := utils.GetFormatFile(newFile)
	path := utils.GetPathFileByFormat(format, newFile.Filename)

	err = g.SaveUploadedFile(newFile, path)

	if err != nil {
		g.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}

	filePath, err := f.fileUC.UploadFile(models.CreateFile{
		Format:   format,
		Path:     path,
		FileName: newFile.Filename,
		UserID:   uuid.NullUUID{},
		Size:     newFile.Size,
	})
	if err != nil {
		g.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}
	g.JSON(http.StatusOK, models.Message{
		Message: filePath,
	})
}

func (f fileHandlers) GetFile(g *gin.Context) {
	fileName := g.Param("file_name")
	foundFile, err := f.fileUC.FindFile(fileName)

	if err != nil {
		g.JSON(http.StatusBadRequest, models.ErrorModel{Error: "file not found"})
		return
	}
	g.File(foundFile.Path)
}

func (f fileHandlers) DeleteFile(g *gin.Context) {
	fileName := g.Param("file_name")
	err := f.fileUC.DeleteFile(fileName)

	if err != nil {
		g.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}

	g.JSON(http.StatusOK, models.Message{Message: "File deleted"})
}

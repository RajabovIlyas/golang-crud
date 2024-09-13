package http

import (
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/file"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/httpErrors"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/httpResponse"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}

	format := utils.GetFormatFile(newFile)
	path := utils.GetPathFileByFormat(format, newFile.Filename)

	err = g.SaveUploadedFile(newFile, path)

	if err != nil {
		g.JSON(httpErrors.ErrorResponse(err))
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
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}
	g.JSON(httpResponse.CreatedResponse(filePath))
}

func (f fileHandlers) GetFile(g *gin.Context) {
	fileName := g.Param("file_name")
	foundFile, err := f.fileUC.FindFile(fileName)

	if err != nil {
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}
	g.File(foundFile.Path)
}

func (f fileHandlers) DeleteFile(g *gin.Context) {
	fileName := g.Param("file_name")
	err := f.fileUC.DeleteFile(fileName)

	if err != nil {
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}

	g.JSON(httpResponse.NoContentResponse("File deleted"))
}

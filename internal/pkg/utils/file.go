package utils

import (
	"fmt"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"mime/multipart"
)

func GetFormatFile(file *multipart.FileHeader) models.Formats {
	switch file.Header.Get("Content-Type") {
	case "image/jpeg", "image/png", "image/gif":
		return models.FormatsPhoto
	case "video/mp4", "video/quicktime", "video/x-msvideo", "video/x-flv":
		return models.FormatsVideo
	case "audio/mpeg", "audio/x-wav", "audio/ogg", "audio/aac":
		return models.FormatsMusic
	default:
		return models.FormatsOther
	}
}

func GetPathFileByFormat(format models.Formats, fileName string) string {
	return fmt.Sprintf("./uploads/%v/%s", format, fileName)
}

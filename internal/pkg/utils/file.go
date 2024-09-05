package utils

import (
	"fmt"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"mime/multipart"
)

func GetFormatFile(file *multipart.FileHeader) database.Formats {
	switch file.Header.Get("Content-Type") {
	case "image/jpeg", "image/png", "image/gif":
		return database.FormatsPhoto
	case "video/mp4", "video/quicktime", "video/x-msvideo", "video/x-flv":
		return database.FormatsVideo
	case "audio/mpeg", "audio/x-wav", "audio/ogg", "audio/aac":
		return database.FormatsMusic
	default:
		return database.FormatsOther
	}
}

func GetPathFileByFormat(format database.Formats, fileName string) string {
	return fmt.Sprintf("./uploads/%v/%s", format, fileName)
}

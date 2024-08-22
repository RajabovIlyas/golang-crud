package common

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
	logger "github.com/rs/zerolog/log"
	"os"
	"time"
)

type FilteredWriter struct {
	w     zerolog.LevelWriter
	level zerolog.Level
}

func (w *FilteredWriter) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}
func (w *FilteredWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level >= w.level {
		return w.w.WriteLevel(level, p)
	}
	return len(p), nil
}

func Logger() {
	_ = os.Mkdir("log", 0777)

	_ = os.Mkdir("log/error", 0777)

	_ = os.Mkdir("log/info", 0777)

	infoFile, err := os.OpenFile(
		"log/info/info-"+time.Now().Format(constants.FormatDate)+".log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		logger.Warn().Msg(err.Error())
	}

	errorFile, err := os.OpenFile(
		"log/error/error-"+time.Now().Format(constants.FormatDate)+".log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Warn(err.Error())
	}

	errWriter := zerolog.MultiLevelWriter(errorFile)
	filteredWriter := &FilteredWriter{errWriter, zerolog.WarnLevel}
	w := zerolog.MultiLevelWriter(infoFile, zerolog.ConsoleWriter{Out: os.Stdout}, filteredWriter)
	logger.Logger = zerolog.New(w).
		With().
		Timestamp().
		Logger()

}

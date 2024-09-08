package logger

import (
	"fmt"
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/rs/zerolog"
	"os"
	"time"
)

const (
	LOG_FOLDER    = "logs"
	ERRORS_FOLDER = "logs/errors"
	INFO_FOLDER   = "logs/info"
	ERROR         = "error"
	INFO          = "info"
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

func getLogFileName(path string, logType string, time string) string {
	return fmt.Sprintf("%s/%s-%v.log", path, logType, time)
}

func InitLogger() (zerolog.Logger, error) {
	_ = os.Mkdir(LOG_FOLDER, 0777)

	_ = os.Mkdir(ERRORS_FOLDER, 0777)

	_ = os.Mkdir(INFO_FOLDER, 0777)

	infoFile, err := os.OpenFile(
		getLogFileName(INFO_FOLDER, INFO, time.Now().Format(constants.FormatDate)),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return zerolog.Logger{}, err
	}

	errorFile, err := os.OpenFile(
		getLogFileName(ERRORS_FOLDER, ERROR, time.Now().Format(constants.FormatDate)),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return zerolog.Logger{}, err
	}

	errWriter := zerolog.MultiLevelWriter(errorFile)
	filteredWriter := &FilteredWriter{errWriter, zerolog.WarnLevel}
	w := zerolog.MultiLevelWriter(infoFile, zerolog.ConsoleWriter{Out: os.Stdout}, filteredWriter)
	logger := zerolog.New(w).
		With().
		Timestamp().
		Logger()
	return logger, nil
}

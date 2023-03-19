package logger

import (
	"fmt"
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func FileLogger(filename string) (logger *zap.Logger, close func()) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	now := time.Now()

	dir, _ := os.Getwd()
	fullPath := dir + "/" + filename
	os.MkdirAll(fullPath, 0755)

	logfile := path.Join(fullPath, fmt.Sprintf("%s.log", now.Format("2006-01-02")))
	logFile, _ := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	close = func() {
		logger.Sync()
		logFile.Close()
	}

	return logger, close
}

// func main() {
//    filename := "logs.log"
//    logger := FileLogger(filename)

//    logger.Info("INFO log level message")
//    logger.Warn("Warn log level message")
//    logger.Error("Error log level message")

// }

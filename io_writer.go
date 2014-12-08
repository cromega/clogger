package clogger

import (
	"fmt"
	"io"
	"log"
)

type writerLogger struct {
	logger *log.Logger
	baseIo io.Closer
}

func CreateIoWriter(target io.WriteCloser) Logger {
	logger := log.New(target, "", 0)
	return &writerLogger{logger: logger, baseIo: target}
}

func (l *writerLogger) log(severity, message string) {
	l.logger.Printf("%v: %v", severity, message)
}

func (l *writerLogger) Debug(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.log("DEBUG", message)
}

func (l *writerLogger) Info(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.log("INFO", message)
}

func (l *writerLogger) Warning(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.log("WARNING", message)
}

func (l *writerLogger) Error(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.log("ERROR", message)
}

func (l *writerLogger) Fatal(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.log("FATAL", message)
}

func (l *writerLogger) Close() {
	l.baseIo.Close()
}

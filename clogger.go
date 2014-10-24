package clogger

import (
	"fmt"
	"io"
	"time"
)

const (
	Fatal = iota
	Error
	Warning
	Info
	Debug
)

type Logger struct {
	target io.Writer
	Level  int
}

func InitLogger(target io.Writer, level int) *Logger {
	return &Logger{target: target, Level: level}
}

func (logger *Logger) log(severity string, message string, objs ...interface{}) {
	fullMessage := fmt.Sprintf(message, objs...)
	time := time.Now().Format("01/02/2006 15:04:05")
	fmt.Fprintf(logger.target, "%s, %s: %s\n", time, severity, fullMessage)
}

func (logger *Logger) Debug(message string, objs ...interface{}) {
	if logger.Level >= 4 {
		logger.log("D", message, objs...)
	}
}

func (logger *Logger) Info(message string, objs ...interface{}) {
	if logger.Level >= 3 {
		logger.log("I", message, objs...)
	}
}

func (logger *Logger) Warning(message string, objs ...interface{}) {
	if logger.Level >= 2 {
		logger.log("W", message, objs...)
	}
}

func (logger *Logger) Error(message string, objs ...interface{}) {
	if logger.Level >= 1 {
		logger.log("E", message, objs...)
	}
}

func (logger *Logger) Fatal(message string, objs ...interface{}) {
	if logger.Level >= 0 {
		logger.log("F", message, objs...)
	}
}

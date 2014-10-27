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
	targets []io.Writer
	Level   int
}

func InitLogger(level int) Logger {
	return Logger{targets: make([]io.Writer, 0), Level: level}
}

func (logger *Logger) AddTarget(target io.Writer) {
	logger.targets = append(logger.targets, target)
}

func (logger *Logger) log(severity string, message string, objs ...interface{}) {
	fullMessage := fmt.Sprintf(message, objs...)
	time := time.Now().Format("01/02/2006 15:04:05")
	for _, target := range logger.targets {
		fmt.Fprintf(target, "%s, %s: %s\n", time, severity, fullMessage)
	}
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

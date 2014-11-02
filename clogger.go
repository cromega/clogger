package clogger

import (
	"fmt"
)

const (
	Fatal = iota
	Error
	Warning
	Info
	Debug
	timeFormat = "01/02/2006 15:04:05"
)

type CloggerTarget interface {
	Write(severity int, message string)
	Close()
}

type Logger struct {
	targets []CloggerTarget
	Level   int
}

func CreateLogger(level int) Logger {
	return Logger{targets: make([]CloggerTarget, 0), Level: level}
}

func (logger *Logger) AddTarget(target CloggerTarget) {
	logger.targets = append(logger.targets, target)
}

func (logger *Logger) log(severity int, message string, objs ...interface{}) {
	fullMessage := fmt.Sprintf(message, objs...)
	for _, target := range logger.targets {
		target.Write(severity, fullMessage)
	}
}

func (logger *Logger) Close() {
	for _, target := range logger.targets {
		target.Close()
	}
}

func (logger *Logger) Debug(message string, objs ...interface{}) {
	if logger.Level >= 4 {
		logger.log(4, message, objs...)
	}
}

func (logger *Logger) Info(message string, objs ...interface{}) {
	if logger.Level >= 3 {
		logger.log(3, message, objs...)
	}
}

func (logger *Logger) Warning(message string, objs ...interface{}) {
	if logger.Level >= 2 {
		logger.log(2, message, objs...)
	}
}

func (logger *Logger) Error(message string, objs ...interface{}) {
	if logger.Level >= 1 {
		logger.log(1, message, objs...)
	}
}

func (logger *Logger) Fatal(message string, objs ...interface{}) {
	if logger.Level >= 0 {
		logger.log(0, message, objs...)
	}
}

package clogger

import (
	"fmt"
	"os"
	"time"
)

type FileTarget struct {
	target *os.File
}

func CreateFileTarget(fileName string) *WriterTarget {
	target, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	return &WriterTarget{target: target}
}

func (self *FileTarget) Write(severity int, message string) {
	var s string
	switch severity {
	case Debug:
		s = "D"
	case Info:
		s = "I"
	case Warning:
		s = "W"
	case Error:
		s = "E"
	case Fatal:
		s = "F"
	}

	t := time.Now().Format(timeFormat)
	message = fmt.Sprintf("%v, %v: %v\n", t, s, message)

	self.target.WriteString(message)
}

func (self *FileTarget) Close() {
	self.target.Close()
}

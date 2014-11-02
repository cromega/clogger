package clogger

import (
	"fmt"
	"io"
	"time"
)

type WriterTarget struct {
	target io.Writer
}

const timeFormat = "01/02/2006 15:04:05"

func CreateWriterTarget(target io.Writer) *WriterTarget {
	return &WriterTarget{target: target}
}

func (self *WriterTarget) Write(severity int, message string) {
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

	fmt.Fprintf(self.target, "%v, %v: %v", t, s, message)
}

func (self *WriterTarget) Close() {
}

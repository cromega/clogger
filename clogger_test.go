package clogger_test

import (
	. "github.com/cromega/clogger"

	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testTarget struct {
	Buffer []string
}

func (self *testTarget) Write(severity int, message string) {
	msg := fmt.Sprintf("%v: %v", severity, message)
	self.Buffer = append(self.Buffer, msg)
	return
}

func (self *testTarget) Close() {
}

var _ = Describe("Clogger", func() {
	var log Logger
	var target testTarget

	BeforeEach(func() {
		target = testTarget{Buffer: make([]string, 0)}
		log = CreateLogger(Debug)
		log.AddTarget(&target)
	})

	It("does debug logging", func() {
		log.Level = Debug
		log.Debug("message %v", 123)

		Expect(target.Buffer[0]).To(Equal("4: message 123"))
	})

	It("does info logging", func() {
		log.Level = Info
		log.Info("message %v", 123)

		Expect(target.Buffer[0]).To(Equal("3: message 123"))
	})

	It("does warning logging", func() {
		log.Level = Warning
		log.Warning("message %v", 123)

		Expect(target.Buffer[0]).To(Equal("2: message 123"))
	})

	It("does error logging", func() {
		log.Level = Error
		log.Error("message %v", 123)

		Expect(target.Buffer[0]).To(Equal("1: message 123"))
	})

	It("does fatal logging", func() {
		log.Level = Fatal
		log.Fatal("message %v", 123)

		Expect(target.Buffer[0]).To(Equal("0: message 123"))
	})

	It("does not log when the severity is low", func() {
		log.Level = Info
		log.Debug("message %v", 123)

		Expect(target.Buffer).To(BeEmpty())
	})

	It("logs to all drains", func() {
		log.Level = Debug
		target2 := testTarget{Buffer: make([]string, 0)}
		log.AddTarget(&target2)
		log.Debug("message %v", 123)

		Expect(len(target.Buffer)).To(Equal(1))
		Expect(len(target2.Buffer)).To(Equal(1))
	})
})

package clogger_test

import (
	. "github.com/cromega/clogger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

// test log drain
type drain struct {
	Buffer []string
}

func (self *drain) Write(data []byte) (n int, err error) {
	str := string(data)
	self.Buffer = append(self.Buffer, str)
	return len(data), nil
}

//

var _ = Describe("Clogger", func() {
	var log Logger
	var target drain

	BeforeEach(func() {
		target = drain{Buffer: make([]string, 0)}
		log = InitLogger(Debug)
		log.AddTarget(&target)
	})

	It("does debug logging", func() {
		log.Level = Debug
		log.Debug("message %v", 123)
		message := strings.SplitN(target.Buffer[0], ", ", 2)[1]

		Expect(message).To(Equal("D: message 123\n"))
	})

	It("does info logging", func() {
		log.Level = Info
		log.Info("message %v", 123)
		message := strings.SplitN(target.Buffer[0], ", ", 2)[1]

		Expect(message).To(Equal("I: message 123\n"))
	})

	It("does warning logging", func() {
		log.Level = Warning
		log.Warning("message %v", 123)
		message := strings.SplitN(target.Buffer[0], ", ", 2)[1]

		Expect(message).To(Equal("W: message 123\n"))
	})

	It("does error logging", func() {
		log.Level = Error
		log.Error("message %v", 123)
		message := strings.SplitN(target.Buffer[0], ", ", 2)[1]

		Expect(message).To(Equal("E: message 123\n"))
	})

	It("does fatal logging", func() {
		log.Level = Fatal
		log.Fatal("message %v", 123)
		message := strings.SplitN(target.Buffer[0], ", ", 2)[1]

		Expect(message).To(Equal("F: message 123\n"))
	})

	It("does not log when the severity is low", func() {
		log.Level = Info
		log.Debug("message %v", 123)

		Expect(target.Buffer).To(BeEmpty())
	})

	It("logs to all drains", func() {
		log.Level = Debug
		target2 := drain{Buffer: make([]string, 0)}
		log.AddTarget(&target2)
		log.Debug("message %v", 123)

		Expect(len(target.Buffer)).To(Equal(1))
		Expect(len(target2.Buffer)).To(Equal(1))
	})
})

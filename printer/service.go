package printer

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type printer struct {
	logger *logrus.Entry
}

func NewPrinter(le *logrus.Entry) Printer{
	return &printer{
		logger: le,
	}
}

func(p printer) Print(message string, data interface{}) {
	p.logger.WithFields(logrus.Fields{
		"data": fmt.Sprintf("%v",data),
	}).Info(message)
}
package main

import (
	"github.com/avasapollo/unravelin/data"
	"github.com/avasapollo/unravelin/encoder"
	"github.com/avasapollo/unravelin/printer"
	"github.com/avasapollo/unravelin/server"
	"github.com/sirupsen/logrus"
)

func main() {
	printerSvc := printer.NewPrinter(logrus.New().WithField("service", "unravelin-api"))
	parserSvc := data.NewParser()
	encoderSvc := encoder.NewHashEncoder()

	server.NewApiRest(printerSvc, parserSvc, encoderSvc).ListenServe(8080)
}

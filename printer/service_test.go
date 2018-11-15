package printer

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

var (
	svc Printer
)

func TestMain(m *testing.M) {
	svc = NewPrinter(logrus.New().WithField("service", "testing"))
	os.Exit(m.Run())
}

func TestPrinter_Print(t *testing.T) {
	var table = []struct {
		name    string
		message string
		data    interface{}
	}{
		{
			name:    "int",
			message: "this interface is an int",
			data:    10,
		},
		{
			name:    "string",
			message: "this interface is a string",
			data:    "hello, my name is Andrea",
		},
		{
			name:    "int",
			message: "this interface is a struct",
			data: struct {
				Name    string
				Surname string
				Age     int
			}{
				Name:    "Andrea",
				Surname: "Vasapollo",
				Age:     30,
			},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			svc.Print(tt.message, tt.data)
		})
	}
}

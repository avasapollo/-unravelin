package printer

type Printer interface {
	Print(message string, data interface{})
}


package clui

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
)

type ConsoleWriter struct{}

func MakeConsoleWriter() ConsoleWriter {
	return ConsoleWriter{}
}

func (printer ConsoleWriter) Write(text string) {
	fmt.Println(text)
}

func (printer ConsoleWriter) WriteWarning(text string) {
	fmt.Println(Red(text))
}
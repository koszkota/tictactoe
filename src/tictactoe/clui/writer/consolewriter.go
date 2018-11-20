package writer

import (
	"fmt"
)

type ConsoleWriter struct{}

func (printer ConsoleWriter) Write(text string) {
	fmt.Println(text)
}

func (printer ConsoleWriter) WriteWarning(text string) {
	fmt.Println("\r\033[31m" + text + "\033[m")
}

func (printer ConsoleWriter) WriteSpinner(text string, spinner string) {
	fmt.Printf(text, spinner)
}
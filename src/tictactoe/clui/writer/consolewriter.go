package writer

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
)

type ConsoleWriter struct{}

func (printer ConsoleWriter) Write(text string) {
	fmt.Println(text)
}

func (printer ConsoleWriter) WriteWarning(text string) {
	fmt.Println(Red(text))
}
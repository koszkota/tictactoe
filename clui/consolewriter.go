package clui

import "fmt"

type ConsoleWriter struct{}

func MakeConsoleWriter() ConsoleWriter {
	return ConsoleWriter{}
}

func (printer ConsoleWriter) Write(text string) {
	fmt.Println(text)
}
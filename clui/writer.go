package clui

import "fmt"

type Writer struct{}

func MakeWriter() Writer {
	return Writer{}
}

func (printer Writer) Write(text string) {
	fmt.Println(text)
}
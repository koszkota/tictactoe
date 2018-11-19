package writer

type Writer interface {
	Write(text string)
	WriteWarning(text string)
}
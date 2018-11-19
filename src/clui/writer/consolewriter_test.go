package writer

func ExampleWritesToConsole() {
	consoleWriter := ConsoleWriter{}
	consoleWriter.Write("test")

	// Output: test
}

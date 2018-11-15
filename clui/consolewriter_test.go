package clui

func ExampleWritesToConsole() {
	consoleWriter := new (ConsoleWriter)
	consoleWriter.Write("test")

	// Output: test
}
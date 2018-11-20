package writer

func ExampleWritesToConsole() {
	consoleWriter := ConsoleWriter{}
	consoleWriter.Write("test")

	// Output: test
}

func ExampleWritesSpinnerToConsole() {
	consoleWriter := ConsoleWriter{}
	consoleWriter.WriteSpinner("test", "...")

	// Output: test%!(EXTRA string=...)
}
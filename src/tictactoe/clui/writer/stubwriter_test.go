package writer

import (
	"testing"
	"tictactoe/testhelper"
)

func TestStubWriterCollectsOutputs(t *testing.T) {
	stubWriter := StubWriter{}
	stubWriter.Write("testWord")
	stubWriter.Write("anotherTestWord")
	expected := [2]string{"testWord", "anotherTestWord"}

	matchers.EqualLiterals(t, expected[0], stubWriter.arrayOfOutputs[0])
	matchers.EqualLiterals(t, expected[1], stubWriter.arrayOfOutputs[1])
}

func TestStubWriterCleansOutputs(t *testing.T) {
	stubWriter := StubWriter{}
	stubWriter.Write("testWord")
	matchers.EqualLiterals(t, "testWord", stubWriter.arrayOfOutputs[0])

	stubWriter.CleanOutputs()

	matchers.EqualLiterals(t, 0, len(stubWriter.arrayOfOutputs))
}

func TestStubWriterReturnsOutput(t *testing.T) {
	stubWriter := StubWriter{}
	stubWriter.Write("testWord")

	matchers.EqualLiterals(t, "testWord", stubWriter.GetOutputs()[0])
}

func TestStubWriterReturnsLastMessage(t *testing.T) {
	stubWriter := StubWriter{}
	stubWriter.Write("testWord")
	stubWriter.Write("anotherTestWord")

	matchers.EqualLiterals(t, "anotherTestWord", stubWriter.GetLastMessage())
}

func TestStubWriterReturnsLastSpinnerMessage(t *testing.T) {
	stubWriter := StubWriter{}
	stubWriter.WriteSpinner("processing", "...")

	matchers.EqualLiterals(t, "processing", stubWriter.GetLastMessage())
}
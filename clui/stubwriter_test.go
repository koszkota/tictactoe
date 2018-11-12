package clui

import (
	"testing"
	"tictactoe/testhelper"
)

func TestStubWriterCollectsOutputs(t *testing.T) {
	stubWriter := StubWriter{}
	stubWriter.Write("koteczek")
	stubWriter.Write("myszek")
	expected := [2]string{"koteczek", "myszek"}
	matchers.EqualLiterals(t, expected[0], stubWriter.arrayOfOutputs[0])
	matchers.EqualLiterals(t, expected[1], stubWriter.arrayOfOutputs[1])
}

func TestStubWriterCleansOutputs(t *testing.T) {
	stubWriter := StubWriter{}
	stubWriter.Write("koteczek")
	matchers.EqualLiterals(t, "koteczek", stubWriter.arrayOfOutputs[0])
	stubWriter.CleanOutputs()
	matchers.EqualLiterals(t, 0, len(stubWriter.arrayOfOutputs))
}

func TestStubWriterReturnsOutpurs(t *testing.T) {
	stubWriter := StubWriter{}
	stubWriter.Write("koteczek")
	matchers.EqualLiterals(t, "koteczek", stubWriter.GetOutputs()[0])
	stubWriter.CleanOutputs()
	matchers.EqualLiterals(t, 0, len(stubWriter.arrayOfOutputs))
}
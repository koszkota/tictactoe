package matchers

import (
	"fmt"
	"testing"
)

func EqualLiterals(tb testing.TB, expected, actual interface{}) {
	if expected != actual {
		fmt.Printf("Equality test failed. %q and %q are not equal", expected, actual)
		tb.FailNow()
	}
}
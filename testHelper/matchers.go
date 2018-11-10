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

func IsTrue(tb testing.TB, actualBoolean bool) {
	fmt.Println(actualBoolean)
	if actualBoolean != true {
		fmt.Printf("The statement is not true")
		tb.FailNow()
	}
}

func IsFalse(tb testing.TB, actualBoolean bool) {
	if actualBoolean != false {
		fmt.Printf("The statement is not false")
		tb.FailNow()
	}
}
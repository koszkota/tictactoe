package controllerstatus

import (
	"testing"
	"tictactoe/src/tictactoe/testhelper"
)

func TestReturnsTrueAsTheFirstValueAndFalseAsSecond(t *testing.T) {
	stubRunStatus := StubRunStatus{0}
	booleanOne := stubRunStatus.GetRunStatus()

	matchers.IsTrue(t, booleanOne)

	booleanTwo := stubRunStatus.GetRunStatus()

	matchers.IsFalse(t, booleanTwo)
}
package controllerstatus

import (
	"testing"
	"tictactoe/testhelper"
)

func TestReturnsTrueAsTheFirstValueAndFalsAsSecojd(t *testing.T) {
	stubRunStatus := StubRunStatus{0}
	booleanOne := stubRunStatus.GetRunStatus()

	matchers.IsTrue(t, booleanOne)

	booleanTwo := stubRunStatus.GetRunStatus()

	matchers.IsFalse(t, booleanTwo)
}
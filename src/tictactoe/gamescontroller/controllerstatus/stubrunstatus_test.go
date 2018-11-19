package controllerstatus

import (
	"testing"
	"../../testhelper"
)

func TestReturnsTrueAsTheFirstValueAndFalseAsSecond(t *testing.T) {
	stubRunStatus := StubRunStatus{0}
	booleanOne := stubRunStatus.GetRunStatus()

	matchers.IsTrue(t, booleanOne)

	booleanTwo := stubRunStatus.GetRunStatus()

	matchers.IsFalse(t, booleanTwo)
}
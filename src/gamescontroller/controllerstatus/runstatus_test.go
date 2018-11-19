package controllerstatus

import (
	"testing"
	"tictactoe/src/testhelper"
)

func TestGetRunStatusReturnsTrue(t *testing.T) {
	runStatus := RunStatus{}
	matchers.IsTrue(t, runStatus.GetRunStatus())
}
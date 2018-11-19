package controllerstatus

import (
	"testing"
	"tictactoe/src/tictactoe/testhelper"
)

func TestGetRunStatusReturnsTrue(t *testing.T) {
	runStatus := RunStatus{}
	matchers.IsTrue(t, runStatus.GetRunStatus())
}
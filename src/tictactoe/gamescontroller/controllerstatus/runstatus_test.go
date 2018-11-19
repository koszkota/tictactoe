package controllerstatus

import (
	"testing"
	"../../testhelper"
)

func TestGetRunStatusReturnsTrue(t *testing.T) {
	runStatus := RunStatus{}
	matchers.IsTrue(t, runStatus.GetRunStatus())
}
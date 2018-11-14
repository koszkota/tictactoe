package gamescontroller

type StubRunStatus struct {
	Counter int
}

func (stubRunStatus *StubRunStatus) GetRunStatus() bool {
	var runStatusBooleans = []bool{true, false}
	var firstBoolean = runStatusBooleans[stubRunStatus.Counter]
	stubRunStatus.Counter += 1
	return firstBoolean
}
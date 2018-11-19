package controllerstatus

type RunStatus struct {}

func (runStatus *RunStatus) GetRunStatus() bool {
	return true
}
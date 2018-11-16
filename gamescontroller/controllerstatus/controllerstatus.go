package controllerstatus

type ControllerStatus interface {
	GetRunStatus() bool
}
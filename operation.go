package retry

//go:generate counterfeiter -o fakes/fake_operation.go . Operation
type Operation interface {
	Try() (retry bool, err error)
}

package retry

//go:generate counterfeiter -o fakes/fake_operation.go . Operation
type Operation interface {
	Try() error
}

//go:generate counterfeiter -o fakes/fake_stopper.go . Stopper
type Stopper interface {
	Stop(err error) bool
}

type Retrier struct {
	operation Operation
	stopper   Stopper
	tries     int
}

func NewRetrier(operation Operation, stopper Stopper) Retrier {
	return Retrier{
		operation: operation,
		stopper:   stopper,
	}
}

func (r *Retrier) Run() (Operation, error) {
	for {
		r.tries++
		err := r.operation.Try()
		if err == nil || r.stopper.Stop(err) {
			return r.operation, err
		}
	}
}

func (r Retrier) Tries() int {
	return r.tries
}

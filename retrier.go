package retry

//go:generate counterfeiter -o fakes/fake_operation.go . Operation
type Operation interface {
	Try() (retry bool, err error)
}

type Retrier struct {
	operation Operation
	tries     int
}

func NewRetrier(operation Operation) Retrier {
	return Retrier{
		operation: operation,
	}
}

func (r *Retrier) Run() (Operation, error) {
	for {
		r.tries++
		retry, err := r.operation.Try()
		if err != nil && retry {
			continue
		}
		return r.operation, err
	}
}

func (r Retrier) Tries() int {
	return r.tries
}

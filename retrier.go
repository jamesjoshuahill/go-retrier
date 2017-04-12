package retry

//go:generate counterfeiter -o fakes/fake_operation.go . Operation
type Operation interface {
	Try() error
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
		err := r.operation.Try()
		if err != nil && isTemporary(err) {
			continue
		}
		return r.operation, err
	}
}

func (r Retrier) Tries() int {
	return r.tries
}

type temporary interface {
	Temporary() bool
}

func isTemporary(err error) bool {
	t, ok := err.(temporary)
	return ok && t.Temporary()
}

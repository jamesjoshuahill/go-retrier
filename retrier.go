package retry

//go:generate counterfeiter -o fakes/fake_operation.go . Operation
type Operation interface {
	Try() error
}

//go:generate counterfeiter -o fakes/fake_failer.go . Failer
type Failer interface {
	Fail(err error) bool
}

type Retrier struct {
	operation Operation
	failer    Failer
	tries     int
}

func NewRetrier(operation Operation, failer Failer) Retrier {
	return Retrier{
		operation: operation,
		failer:    failer,
	}
}

func (r *Retrier) Run() (Operation, error) {
	for {
		r.tries++
		err := r.operation.Try()
		if err == nil || r.failer.Fail(err) {
			return r.operation, err
		}
	}
}

func (r Retrier) Tries() int {
	return r.tries
}

package retry

type Retrier struct {
	operation Operation
	tries     int
}

func NewRetrier(operation Operation) Retrier {
	return Retrier{
		operation: operation,
	}
}

func (r *Retrier) Run() error {
	for {
		r.tries++
		r.operation.Try()
		if !r.operation.Retry() {
			break
		}
	}

	return r.operation.Error()
}

func (r Retrier) Tries() int {
	return r.tries
}

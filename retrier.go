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
	r.tries++
	return r.operation.Try()
}

func (r Retrier) Tries() int {
	return r.tries
}

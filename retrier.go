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
		retry, err := r.operation.Try()
		if err == nil {
			return nil
		}
		if !retry {
			return err
		}
	}
}

func (r Retrier) Tries() int {
	return r.tries
}

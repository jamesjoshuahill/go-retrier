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
	retryable, err := r.operation.Try()
	if err != nil {
		if !retryable {
			return err
		}
		return r.Run()
	}

	return nil
}

func (r Retrier) Tries() int {
	return r.tries
}

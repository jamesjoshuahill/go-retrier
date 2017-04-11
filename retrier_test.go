package retry_test

import (
	"errors"

	"github.com/jamesjoshuahill/go-retry"
	"github.com/jamesjoshuahill/go-retry/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Retrier", func() {
	Context("when the operation succeeds", func() {
		It("tries once", func() {
			operation := new(fakes.FakeOperation)
			operation.RetryReturns(false)
			operation.ErrorReturns(nil)
			retrier := retry.NewRetrier(operation)

			err := retrier.Run()

			Expect(err).NotTo(HaveOccurred())
			Expect(retrier.Tries()).To(Equal(1))
			Expect(operation.TryCallCount()).To(Equal(1))
		})
	})

	Context("when the operation fails", func() {
		Context("and is not retryable", func() {
			It("returns the error", func() {
				operation := new(fakes.FakeOperation)
				operationErr := errors.New("operation failed")
				operation.RetryReturns(false)
				operation.ErrorReturns(operationErr)
				retrier := retry.NewRetrier(operation)

				err := retrier.Run()

				Expect(err).To(Equal(operationErr))
				Expect(retrier.Tries()).To(Equal(1))
				Expect(operation.TryCallCount()).To(Equal(1))
			})
		})

		Context("and is retryable", func() {
			It("retries the operation", func() {
				operation := new(fakes.FakeOperation)
				operation.RetryReturnsOnCall(0, true)
				operation.RetryReturnsOnCall(1, false)
				operation.ErrorReturns(nil)
				retrier := retry.NewRetrier(operation)

				err := retrier.Run()

				Expect(err).NotTo(HaveOccurred())
				Expect(retrier.Tries()).To(Equal(2))
				Expect(operation.TryCallCount()).To(Equal(2))
			})
		})
	})

	Context("example operation", func() {
		It("tries once", func() {
			operation := NewSimpleOperation()
			retrier := retry.NewRetrier(&operation)

			err := retrier.Run()

			Expect(err).NotTo(HaveOccurred())
			Expect(retrier.Tries()).To(Equal(2))
		})
	})
})

type simpleOperation struct {
	attempts    int
	maxAttempts int
	retry       bool
	err         error
}

func NewSimpleOperation() simpleOperation {
	return simpleOperation{
		maxAttempts: 2,
	}
}

func (o *simpleOperation) Try() {
	o.attempts++
	if o.attempts == o.maxAttempts {
		o.retry = false
	} else {
		o.retry = true
	}
	o.err = errors.New("blurgh")
}

func (o simpleOperation) Retry() bool {
	return o.retry
}

func (o simpleOperation) Error() error {
	return o.err
}

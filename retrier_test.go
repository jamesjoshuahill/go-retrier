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
})

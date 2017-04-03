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
			operation.TryReturns(nil)
			retrier := retry.NewRetrier(operation)

			err := retrier.Run()

			Expect(err).NotTo(HaveOccurred())
			Expect(retrier.Tries()).To(Equal(1))
			Expect(operation.TryCallCount()).To(Equal(1))
		})
	})

	Context("when the operation fails", func() {
		It("returns the error", func() {
			operation := new(fakes.FakeOperation)
			operationErr := errors.New("operation failed")
			operation.TryReturns(operationErr)
			retrier := retry.NewRetrier(operation)

			err := retrier.Run()

			Expect(err).To(Equal(operationErr))
			Expect(retrier.Tries()).To(Equal(1))
			Expect(operation.TryCallCount()).To(Equal(1))
		})
	})
})

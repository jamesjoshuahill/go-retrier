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
		It("returns the operation", func() {
			operation := new(fakes.FakeOperation)
			operation.TryReturns(nil)
			retrier := retry.NewRetrier(operation)

			actualOperation, err := retrier.Run()

			Expect(err).NotTo(HaveOccurred())
			Expect(retrier.Tries()).To(Equal(1))
			Expect(operation.TryCallCount()).To(Equal(1))
			Expect(actualOperation).To(Equal(operation))
		})
	})

	Context("when the operation fails", func() {
		It("returns the error", func() {
			operation := new(fakes.FakeOperation)
			operationErr := errors.New("operation failed")
			operation.TryReturns(operationErr)
			retrier := retry.NewRetrier(operation)

			actualOperation, err := retrier.Run()

			Expect(err).To(Equal(operationErr))
			Expect(retrier.Tries()).To(Equal(1))
			Expect(operation.TryCallCount()).To(Equal(1))
			Expect(actualOperation).To(Equal(operation))
		})
	})

	Context("when the operation has a temporary error", func() {
		It("trys again", func() {
			operation := new(fakes.FakeOperation)
			operation.TryReturnsOnCall(0, temporaryError{})
			operation.TryReturnsOnCall(1, nil)
			retrier := retry.NewRetrier(operation)

			actualOperation, err := retrier.Run()

			Expect(err).NotTo(HaveOccurred())
			Expect(retrier.Tries()).To(Equal(2))
			Expect(operation.TryCallCount()).To(Equal(2))
			Expect(actualOperation).To(Equal(operation))
		})
	})
})

type temporaryError struct{}

func (e temporaryError) Error() string   { return "temporary error" }
func (e temporaryError) Temporary() bool { return true }

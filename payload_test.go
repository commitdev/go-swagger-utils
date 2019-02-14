package utils_test

import (
	"errors"

	. "github.com/forksu/go-swagger-utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Payload", func() {
	Describe("GetErrorPayload", func() {
		It("should return error payload correctly", func() {
			message := "Error message"
			code := 500
			err := errors.New(message)
			payload := GetErrorPayload(code, err)

			Expect(payload.Code).To(Equal(int64(code)))
			Expect(*payload.Message).To(Equal(message))
		})
	})
})

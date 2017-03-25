package twiml

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validators", func() {
	It("can validate one of several options", func() {
		notok := Validate(OneOf("test", "this", "that"))
		ok := Validate(OneOf("test", "test", "that"))
		Expect(notok).To(Equal(false))
		Expect(ok).To(Equal(true))
	})

	It("can validate several validators at once", func() {
		ok := Validate(
			OneOf("test", "test"),
			OneOf("foo", "bar", "baz"),
		)
		Expect(ok).To(Equal(false))
	})
})

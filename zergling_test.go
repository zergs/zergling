package zergling_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Zergling", func() {
	It("should pass", func(){
		Expect("zergling").To(Equal("zergling"))
	})
})

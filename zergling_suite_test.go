package zergling_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestZergling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Zergling Suite")
}

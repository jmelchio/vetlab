package vetlab_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVetlab(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vetlab Suite")
}

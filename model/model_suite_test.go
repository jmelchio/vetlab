package model_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Model Suite")
}

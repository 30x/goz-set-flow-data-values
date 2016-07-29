package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGozSetFlowDataValues(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GozSetFlowDataValues Suite")
}

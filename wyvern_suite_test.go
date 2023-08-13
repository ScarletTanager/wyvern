package wyvern_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWyvern(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wyvern Suite")
}

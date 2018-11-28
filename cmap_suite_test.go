package cmap_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cmap suite")
}

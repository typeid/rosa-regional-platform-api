package e2e_cli_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestE2ECLI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ROSA Regional Platform API CLI E2E Suite")
}

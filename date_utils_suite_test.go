package date_utils_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDateUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DateUtils Suite")
}

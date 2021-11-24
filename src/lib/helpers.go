package lib

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/tebeka/selenium"
	"io/ioutil"
	"log"
	"time"
)

var (
	Fail        = ginkgo.Fail
	RunSpecs    = ginkgo.RunSpecs
	Describe    = ginkgo.Describe
	BeforeEach  = ginkgo.BeforeEach
	AfterEach   = ginkgo.AfterEach
	It          = ginkgo.It
	CurrentTest = ginkgo.CurrentGinkgoTestDescription

	RegisterFailHandler = gomega.RegisterFailHandler
	Expect              = gomega.Expect
	HaveOccurred        = gomega.HaveOccurred
	BeZero              = gomega.BeZero
)

// URL returns full path for passed environment value.
func URL(env Env) string { return urlMap[env] }

// TakeScreenshot saves screenshot of passed WebDriver into file with passed test name.
func TakeScreenshot(wd selenium.WebDriver, testName string) {
	bytes, err := wd.Screenshot()
	if err != nil {
		log.Panic("Can't take a screenshot.")
	}
	ioutil.WriteFile(testName+".jpg", bytes, 0644)
}

func ErrCheck(err error) {
	Expect(err).ToNot(HaveOccurred())
}

// MustNotFindElement returns fails if element is found.
func MustNotFindElement(wd selenium.WebDriver, by, value string) {
	wd.SetImplicitWaitTimeout(time.Second)
	defer wd.SetImplicitWaitTimeout(DefTimeout)
	element, err := wd.FindElement(by, value)
	Expect(element).To(BeZero())
	Expect(err).To(HaveOccurred())
}


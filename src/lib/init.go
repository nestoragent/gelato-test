package lib

import (
	"fmt"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/tebeka/selenium"
	"io/ioutil"
	"log"
	"time"
)

type Env string
type Browser string

const (
	DevEnv     Env     = "dev"
	Chrome     Browser = "chrome"
	DefTimeout         = 5 * time.Second
)

var UrlMap = map[Env]string{
	DevEnv: "https://todomvc.com/examples/react/#/",
}

// Conf represents configuration data.
type Conf struct {
	Browser  Browser
	Env      Env
	Headless bool
	Port     int
	Width    int
	Height   int
}

var (
	conf                Conf
	caps                selenium.Capabilities
	Expect              = gomega.Expect
	HaveOccurred        = gomega.HaveOccurred
	BeZero              = gomega.BeZero
	RegisterFailHandler = gomega.RegisterFailHandler
	Fail                = ginkgo.Fail
	RunSpecs            = ginkgo.RunSpecs
)

func GetConf() Conf { return conf }

func SetCaps(cnf Conf) {
	switch cnf.Browser {
	case Chrome:
		setChromeCaps(cnf)
	default:
		log.Panic("Invalid Browser type.")
	}
}

func GetCaps() selenium.Capabilities { return caps }

func setChromeCaps(cnf Conf) {
	args := []string{
		fmt.Sprintf("--window-size=%d,%d", cnf.Width, cnf.Height),
		"--ignore-certificate-errors",
		"--disable-extensions",
		"--no-sandbox",
		"--disable-dev-shm-usage",
	}
	if cnf.Headless {
		args = append(args, "--headless", "--disable-gpu")
	}
	chromeCaps := map[string]interface{}{
		"excludeSwitches": [1]string{"enable-automation"},
		"args":            args,
		"w3c":             false,
	}
	caps = selenium.Capabilities{
		"browserName":    "chrome",
		"chromeOptions":  chromeCaps,
		"browserVersion": "93.0",
		"enableVNC":      true,
		"enableVideo":    true,
	}
}

// URL returns full path for passed environment value.
func URL(env Env) string { return UrlMap[env] }

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

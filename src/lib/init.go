package lib

import (
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	"time"
)

type Env string
type Browser string

const (
	DevEnv Env     = "dev"
	Chrome Browser = "chrome"
	DefTimeout = 5 * time.Second
)

var urlMap = map[Env]string{
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
	conf Conf
	caps selenium.Capabilities
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
		"w3c": false,
	}
	caps = selenium.Capabilities{
		"browserName":    "chrome",
		"chromeOptions":  chromeCaps,
		"browserVersion": "93.0",
		"enableVNC": true,
		"enableVideo": true,
	}
}

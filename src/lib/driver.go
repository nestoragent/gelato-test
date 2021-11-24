package lib

import (
	"fmt"
	"github.com/tebeka/selenium"
)

var (
	driver selenium.WebDriver
)
func GetDriver() selenium.WebDriver {
	if driver == nil {
		driver = InitDriver()
	}
	return driver
}

func InitDriver() selenium.WebDriver {
	//driver, err := selenium.NewRemote(GetCaps(), "http://localhost:4444/wd/hub")
	driver, err := selenium.NewRemote(GetCaps(), fmt.Sprintf("http://localhost:%d/wd/hub", GetConf().Port))

	if err != nil {
		panic(fmt.Sprintf("create selenium session: %v", err))
	}
	ErrCheck(err)
	Expect(driver).ToNot(BeZero())
	ErrCheck(driver.SetImplicitWaitTimeout(DefTimeout))
	return driver
}

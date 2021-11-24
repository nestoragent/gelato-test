package tests

import (
	"fmt"
	. "gelato-test/src/lib"
	"github.com/tebeka/selenium"
	"testing"
)

// User test suite.
func TestUser(t *testing.T) {
	DefaultConf()
	RegisterFailHandler(Fail)
	RunSpecs(t, "User")
}

var _ = Describe("User", func() {
	var (
		driver      selenium.WebDriver

		//username = fmt.Sprintf("batman_%v", time.Now().Unix())
		//password = "secret123"
	)

	BeforeEach(func() {
		var err error
		//service, err = StartSelenium()
		//Expect(service).ToNot(BeZero())
		//ErrCheck(err)
		//ErrCheck(err)
		driver, err = selenium.NewRemote(GetCaps(), "http://localhost:4444/wd/hub")
		//driver, err = selenium.NewRemote(GetCaps(), fmt.Sprintf("http://selenoid:%d/driver/hub", GetConf().Port))

		if err != nil {
			panic(fmt.Sprintf("create selenium session: %v", err))
		}
		ErrCheck(err)
		Expect(driver).ToNot(BeZero())
		ErrCheck(driver.SetImplicitWaitTimeout(DefTimeout))
		ErrCheck(driver.Get(URL(GetConf().Env)))
	})

	AfterEach(func() {
		TakeScreenshot(driver, CurrentTest().TestText)
		err := driver.Quit()
		ErrCheck(err)
	})

	It("can create new account", func() {
		//loginLink := MustFindElement(driver, selenium.ByLinkText, "LOGIN")
		//ErrCheck(loginLink.Click())
		//newAccountLink := MustFindElement(driver, selenium.ByCSSSelector, ".btn-link")
		//ErrCheck(newAccountLink.Click())
		//usernameInput := MustFindElement(driver, selenium.ByID, "username")
		//ErrCheck(usernameInput.SendKeys(username))
		//passwordInput := MustFindElement(driver, selenium.ByID, "password")
		//ErrCheck(passwordInput.SendKeys(password))
		//submitButton := MustFindElement(driver, selenium.ByCSSSelector, ".btn-success")
		//ErrCheck(submitButton.Click())
		//MustWaitWithTimeout(driver, func(driver selenium.WebDriver) (bool, error) {
		//	header := MustFindElement(driver, selenium.ByTagName, "h1")
		//	text, err := header.Text()
		//	return text == "Welcome to React Gin Blog!", err
		//}, 5*time.Second)
		//logoutLink := MustFindElement(driver, selenium.ByCSSSelector, ".btn-dark")
		//Expect(logoutLink).ToNot(BeZero())
	})
})

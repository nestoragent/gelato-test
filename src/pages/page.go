package pages

import (
	"fmt"
	. "gelato-test/src/lib"
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"time"
)

type Page struct {
	Driver selenium.WebDriver
}

func (s *Page) getDriver() selenium.WebDriver {
	return s.Driver
}

func (s *Page) FindElementById(locator string) selenium.WebElement {
	element, _ := s.Driver.FindElement(selenium.ByID, locator)
	return element
}

func (s *Page) FindElementByXpath(locator string) selenium.WebElement {
	element, _ := s.Driver.FindElement(selenium.ByXPATH, locator)
	return element
}

func (s *Page) FindElementByLinkText(locator string) selenium.WebElement {
	element, _ := s.Driver.FindElement(selenium.ByLinkText, locator)
	return element
}

func (s *Page) FindElementByPartialLink(locator string) selenium.WebElement {
	element, _ := s.Driver.FindElement(selenium.ByPartialLinkText, locator)
	return element
}

func (s *Page) FindElementByName(locator string) selenium.WebElement {
	element, _ := s.Driver.FindElement(selenium.ByName, locator)
	return element
}

func (s *Page) FindElementByTag(locator string) selenium.WebElement {
	element, _ := s.Driver.FindElement(selenium.ByTagName, locator)
	return element
}

func (s *Page) FindElementByClass(locator string) selenium.WebElement {
	element, _ := s.Driver.FindElement(selenium.ByClassName, locator)
	return element
}

func (s *Page) FindElementByCss(locator string) selenium.WebElement {
	element, err := s.Driver.FindElement(selenium.ByCSSSelector, locator)
	ErrCheck(err)
	return element
}

func (s *Page) FindElementsById(locator string) []selenium.WebElement {
	element, _ := s.Driver.FindElements(selenium.ByID, locator)
	return element
}

func (s *Page) FindElementsByXpath(locator string) []selenium.WebElement {
	element, _ := s.Driver.FindElements(selenium.ByXPATH, locator)
	return element
}

func (s *Page) FindElementsByLinkText(locator string) []selenium.WebElement {
	element, _ := s.Driver.FindElements(selenium.ByLinkText, locator)
	return element
}

func (s *Page) FindElementsByPartialLink(locator string) []selenium.WebElement {
	element, _ := s.Driver.FindElements(selenium.ByPartialLinkText, locator)
	return element
}

func (s *Page) FindElementsByName(locator string) []selenium.WebElement {
	element, _ := s.Driver.FindElements(selenium.ByName, locator)
	return element
}

func (s *Page) FindElementsByTag(locator string) []selenium.WebElement {
	element, _ := s.Driver.FindElements(selenium.ByTagName, locator)
	return element
}

func (s *Page) FindElementsByClass(locator string) []selenium.WebElement {
	element, _ := s.Driver.FindElements(selenium.ByClassName, locator)
	return element
}

func (s *Page) FindElementsByCss(locator string) []selenium.WebElement {
	element, _ := s.Driver.FindElements(selenium.ByCSSSelector, locator)
	return element
}

func (s *Page) MouseHoverToElement(locator string) selenium.WebElement {
	element, _ := s.Driver.FindElement(selenium.ByCSSSelector, locator)
	err := element.MoveTo(0, 0)
	if err != nil {
		panic(fmt.Sprintf("Error in the MouseHoverToElement, err: %v", err))
	}
	return element
}

func (s *Page) WaitWithTimeout(condition selenium.Condition, timeout time.Duration) {
	ErrCheck(s.Driver.WaitWithTimeout(condition, timeout))
}


func assertTrue(result bool) error {
	var t asserter
	assert.True(&t, result)

	return t.err
}

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}

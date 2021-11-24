package tests

import (
	"fmt"
	. "gelato-test/src/lib"
	"gelato-test/src/pages"
	"github.com/tebeka/selenium"
	"testing"
)

func TestToDos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ToDos")
}

var _ = Describe("ToDos", func() {
	var (
		driver selenium.WebDriver
		page   pages.Page
	)

	BeforeEach(func() {
		DefaultConf()
		driver = GetDriver()
	})

	AfterEach(func() {
		TakeScreenshot(driver, CurrentTest().TestText)
		err := driver.Quit()
		ErrCheck(err)
	})

	It("can create new todos", func() {
		fmt.Printf("Result url: %v", URL(GetConf().Env))
		ErrCheck(driver.Get(URL(GetConf().Env)))

		page = pages.Page{Driver: driver}
		todos := pages.ToDos{Page: page}
		todos.AddToDo("test todo")
	})
})

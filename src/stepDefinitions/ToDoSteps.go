package stepDefinitions

import (
	"fmt"
	. "gelato-test/src/lib"
	"gelato-test/src/pages"
	"github.com/cucumber/godog"
)

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		fmt.Println("Work BeforeSuite!")
		RegisterFailHandler(Fail)
		DefaultConf()
		//GetDriver()
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	fmt.Println("Work InitializeScenario!")
	//ctx.BeforeScenario(func(*godog.Scenario) {
	//	Godogs = 0 // clean the state before every scenario
	//})

	ctx.Step(`^go to main todos page$`, goToMainPage)
	ctx.Step(`^create a new task "([^"]*)"$`, createANewToDo)
	ctx.Step(`^task "([^"]*)" exist in the list$`, taskExistInTheList)
}

func goToMainPage() {
	ErrCheck(GetDriver().Get(URL(GetConf().Env)))
}

func createANewToDo(name string) {
	var page = pages.Page{Driver: GetDriver()}
	todos := pages.ToDos{Page: page}
	todos.AddToDo(name)
}

func taskExistInTheList(name string) {
	var page = pages.Page{Driver: GetDriver()}
	todos := pages.ToDos{Page: page}
	todos.CheckThatRowExist(name)
}

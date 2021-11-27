package stepDefinitions

import (
	"context"
	. "gelato-test/src/lib"
	"gelato-test/src/pages"
	"github.com/cucumber/godog"
)

var (
	page pages.Page
)

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	//ctx.BeforeSuite(func() {
	//	fmt.Println("Before Suite")
	//})
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		RegisterFailHandler(Fail)
		DefaultConf()
		page = pages.Page{Driver: GetDriver()}
		return ctx, nil
	})

	ctx.Step(`^go to main todos page$`, goToMainPage)
	ctx.Step(`^check that input is empty$`, checkThatInputIsEmpty)
	ctx.Step(`^create a new task "([^"]*)"$`, createANewToDo)
	ctx.Step(`^task "([^"]*)" exist in the list$`, taskExistInTheList)
	ctx.Step(`^task "([^"]*)" not exist in the list$`, taskNotExistInTheList)
	ctx.Step(`^item count should be "([^"]*)"$`, checkItemCount)
	ctx.Step(`^mark item "([^"]*)" as completed$`, markAsCompleted)
	ctx.Step(`^check that item "([^"]*)" is completed$`, checkThatTaskCompleted)
	ctx.Step(`^switch tab to "([^"]*)"$`, switchTab)
	ctx.Step(`^delete item "([^"]*)"$`, deleteItem)
	ctx.Step(`^delete completed items$`, deleteCompletedItem)
	ctx.Step(`^rename item "([^"]*)" to "([^"]*)"$`, renameItem)

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		CloseDriver()
		return ctx, err
	})
}

func goToMainPage() {
	ErrCheck(GetDriver().Get(URL(GetConf().Env)))
}

func createANewToDo(name string) {
	todos := pages.ToDos{Page: page}
	todos.AddToDo(name)
}

func taskExistInTheList(name string) {
	todos := pages.ToDos{Page: page}
	todos.CheckThatRowExist(name)
}

func taskNotExistInTheList(name string) {
	todos := pages.ToDos{Page: page}
	todos.CheckThatItemNotInList(name)
}

func checkThatInputIsEmpty() {
	todos := pages.ToDos{Page: page}
	todos.CheckThatInputIsEmpty()
}

func checkItemCount(count string) {
	todos := pages.ToDos{Page: page}
	todos.CheckItemsCount(count)
}

func markAsCompleted(name string) {
	todos := pages.ToDos{Page: page}
	todos.MarkAsCompleted(name)
}

func checkThatTaskCompleted(name string) {
	todos := pages.ToDos{Page: page}
	todos.CheckThatTaskCompleted(name)
}

func switchTab(tabName string) {
	todos := pages.ToDos{Page: page}
	todos.SwitchTab(tabName)
}

func deleteItem(itemName string) {
	todos := pages.ToDos{Page: page}
	todos.DeleteItem(itemName)
}

func deleteCompletedItem() {
	todos := pages.ToDos{Page: page}
	todos.DeleteCompletedItem()
}

func renameItem(oldName string, newName string) {
	todos := pages.ToDos{Page: page}
	todos.RenameItem(oldName, newName)
}

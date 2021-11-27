package pages

import (
	"fmt"
	. "gelato-test/src/lib"
	"github.com/tebeka/selenium"
	"time"
)

type ToDos struct {
	Page Page
}

var (
	cssInputNewToDo           = "input.new-todo"
	xpathToDoRow              = "//label[text()='%s']/.."
	xpathToDoRowMarkCompleted = "//label[text()='%s']/../input[@class='toggle']"
	xpathToDoRowCompleted     = "//label[text()='%s']/../.."
	cssItemsCount             = "span.todo-count > strong"
	cssTabActive              = "a[href='#/active']"
	cssTabCompleted           = "a[href='#/completed']"
	cssTabAll                 = "a[href='#/']"
	xpathToDoRowDestroy       = "//label[text()='%s']/../button[@class='destroy']"
	xpathToDoLabel            = "//label[text()='%s']"
	cssToDoRowRename          = "li.editing input.edit"
	cssButtonClearCompleted   = "button.clear-completed"
)

func (s *ToDos) AddToDo(toDoName string) *ToDos {
	s.Page.FindElementByCss(cssInputNewToDo).SendKeys(toDoName)
	s.Page.FindElementByCss(cssInputNewToDo).SendKeys(selenium.EnterKey)
	return &ToDos{Page: s.Page}
}

func (s *ToDos) CheckThatRowExist(toDoName string) *ToDos {
	var todos = s.Page.FindElementsByXpath(fmt.Sprintf(xpathToDoRow, toDoName))
	if len(todos) == 0 {
		panic(fmt.Sprintf("Can not find the item. Actual size: %v", len(todos)))
	}
	return &ToDos{Page: s.Page}
}

func (s *ToDos) CheckThatItemNotInList(toDoName string) *ToDos {
	var todos = s.Page.FindElementsByXpath(fmt.Sprintf(xpathToDoRow, toDoName))

	if len(todos) != 0 {
		panic(fmt.Sprintf("Item should not to be in list. Actual size: %v", len(todos)))
	}
	return &ToDos{Page: s.Page}
}

func (s *ToDos) CheckThatInputIsEmpty() *ToDos {
	var value, _ = s.Page.FindElementByCss(cssInputNewToDo).GetAttribute("value")

	if value != "" {
		panic("Input should be empty")
	}
	return &ToDos{Page: s.Page}
}

func (s *ToDos) CheckItemsCount(count string) *ToDos {
	var text, _ = s.Page.FindElementByCss(cssItemsCount).Text()
	if text != count {
		panic("Count is not correct!")
	}
	return &ToDos{Page: s.Page}
}

func (s *ToDos) MarkAsCompleted(name string) *ToDos {
	s.Page.FindElementByXpath(fmt.Sprintf(xpathToDoRowMarkCompleted, name)).Click()
	return &ToDos{Page: s.Page}
}

func (s *ToDos) CheckThatTaskCompleted(name string) *ToDos {
	var class, _ = s.Page.FindElementByXpath(fmt.Sprintf(xpathToDoRowCompleted, name)).GetAttribute("class")
	if class != "completed" {
		panic(fmt.Sprintf("Task %v should be completed!", name))
	}
	return &ToDos{Page: s.Page}
}

func (s *ToDos) SwitchTab(tabName string) *ToDos {
	var tabElem selenium.WebElement
	switch tabName {
	case "All":
		tabElem = s.Page.FindElementByCss(cssTabAll)
	case "Active":
		tabElem = s.Page.FindElementByCss(cssTabActive)
	case "Completed":
		tabElem = s.Page.FindElementByCss(cssTabCompleted)
	default:
		panic(fmt.Sprintf("Unknown tab: %v", tabName))
	}
	tabElem.Click()
	return &ToDos{Page: s.Page}
}

func (s *ToDos) DeleteItem(name string) *ToDos {
	s.Page.FindElementByXpath(fmt.Sprintf(xpathToDoRow, name)).Click()
	s.Page.FindElementByXpath(fmt.Sprintf(xpathToDoRowDestroy, name)).Click()
	return &ToDos{Page: s.Page}
}

func (s *ToDos) DeleteCompletedItem() *ToDos {
	s.Page.FindElementByCss(cssButtonClearCompleted).Click()
	return &ToDos{Page: s.Page}
}

func (s *ToDos) RenameItem(oldName string, newName string) *ToDos {
	var err = s.Page.FindElementByXpath(fmt.Sprintf(xpathToDoLabel, oldName)).MoveTo(10, 10)
	if err != nil {
		panic(err)
	}
	s.Page.FindElementByXpath(fmt.Sprintf(xpathToDoLabel, oldName)).Click()
	err = GetDriver().DoubleClick()
	if err != nil {
		panic(err)
	}

	s.Page.WaitWithTimeout(func(driver selenium.WebDriver) (bool, error) {
		return s.Page.FindElementByCss(cssToDoRowRename).IsEnabled()
	}, 5*time.Second)

	var inputRename = s.Page.FindElementByCss(cssToDoRowRename)
	value, _ := inputRename.GetAttribute("value")

	for i := 0; i < len(value); i++ {
		err = inputRename.SendKeys(selenium.BackspaceKey)
		if err != nil {
			panic(err)
		}
	}

	err = inputRename.SendKeys(newName)
	if err != nil {
		panic(err)
	}
	err = s.Page.FindElementByCss(cssInputNewToDo).Click()
	if err != nil {
		panic(err)
	}
	return &ToDos{Page: s.Page}
}

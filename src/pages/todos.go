package pages

import (
	"fmt"
	"github.com/tebeka/selenium"
)

type ToDos struct {
	Page Page
}

var (
	inputNewToDo = "input.new-todo"
	xpathToDoRow = "//label[text()='%s']/.."
)

func (s *ToDos) AddToDo(toDoName string) *ToDos {
	s.Page.FindElementByCss(inputNewToDo).SendKeys(toDoName)
	s.Page.FindElementByCss(inputNewToDo).SendKeys(selenium.EnterKey)
	return &ToDos{Page:s.Page}
}

func (s *ToDos) CheckThatRowExist(toDoName string) *ToDos {
	fmt.Printf("result path: " + fmt.Sprintf(xpathToDoRow, toDoName))
	var todos = s.Page.FindElementsByXpath(fmt.Sprintf(xpathToDoRow, toDoName))

	if len(todos) == 0 {
		panic(fmt.Sprintf("Can not find the item. Actual size: %v", len(todos)))
	}
	return &ToDos{Page:s.Page}
}


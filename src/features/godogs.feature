Feature: ToDos feature

  @one
  Scenario: Should allow to add todo item
    When go to main todos page
    When create a new task "First Item"
    Then task "First Item" exist in the list
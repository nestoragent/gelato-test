@RT
Feature: ToDos feature

  Background:
    Given go to main todos page

  @one
  Scenario: Should allow to add todo item
    When create a new task "First Item"
    Then task "First Item" exist in the list

  @two
  Scenario Outline: Add items
    When create a new task "<name>"
    Then task "<name>" exist in the list

    Examples:
      | name        |
      | First Item  |
      | Second Item |
      | Thirds Item |

  @three
  Scenario: Should clear text input field when an item is added
    When create a new task "Empty Item"
    Then task "Empty Item" exist in the list
    Then check that input is empty

  @four
  Scenario: Should trim text input
    When create a new task "     Item with spaces      "
    Then task "Item with spaces" exist in the list

  @five
  Scenario: Should shown items count in bottom
    When create a new task "One"
    When create a new task "Two"
    When create a new task "Third"
    When create a new task "Four"
    Then item count should be "4"

  @six
  Scenario: Can mark as completed
    When create a new task "One"
    When create a new task "Two"
    When mark item "One" as completed
    When check that item "One" is completed

  @seven
  Scenario: Can filter items
    When create a new task "One"
    When create a new task "Two"
    When create a new task "Third"
    When create a new task "Four"
    When mark item "One" as completed
    When mark item "Third" as completed
    When switch tab to "Active"
    Then task "One" not exist in the list
    Then task "Two" exist in the list
    Then task "Third" not exist in the list
    Then task "Four" exist in the list
    When switch tab to "Completed"
    Then task "One" exist in the list
    Then task "Two" not exist in the list
    Then task "Third" exist in the list
    Then task "Four" not exist in the list

  @eight
  Scenario: Can delete item
    When create a new task "One"
    When create a new task "Two"
    When create a new task "Third"
    When create a new task "Four"
    When delete item "One"
    Then task "One" not exist in the list

  @nine
  Scenario: Can rename item
    When create a new task "One"
    When create a new task "Two"
    When rename item "One" to "One New"
    Then task "One New" exist in the list

  @ten
  Scenario: Can delete all complete tasks
    When create a new task "One"
    When create a new task "Two"
    When create a new task "Third"
    When create a new task "Four"
    When mark item "One" as completed
    When mark item "Third" as completed
    When delete completed items
    Then task "One" not exist in the list
    Then task "Two" exist in the list
    Then task "Third" not exist in the list
    Then task "Four" exist in the list


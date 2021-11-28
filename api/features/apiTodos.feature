@API
Feature: Api ToDos feature

  @one
  Scenario: Should allow get todo item by id
    When I send "GET" request to "https://jsonplaceholder.typicode.com/todos/1" without body
    Then the response code should be 200
    And the response should match json:
      """
      {
          "userId": 1,
          "id": 1,
          "title": "delectus aut autem",
          "completed": false
      }
      """

  @two
  Scenario: Should allow creating a resource
    When I send "POST" request to "https://jsonplaceholder.typicode.com/posts" with:
      | title | body | userId |
      | foo   | bar  | 1      |
    Then the response code should be 201
    And the response should match json:
      """
      {
          "body": "bar",
          "title": "foo",
          "userId": "1",
          "id": 101
      }
      """


  @third
  Scenario: Should allow updating a resource
    When I send "PUT" request to "https://jsonplaceholder.typicode.com/posts/1" with:
      | id | title | body | userId |
      | 1  | foo1  | bar1 | 11     |
    Then the response code should be 200
    And the response should match json:
      """
      {
          "body": "bar1",
          "title": "foo1",
          "userId": "11",
          "id": 1
      }
      """


  @four
  Scenario: Should allow patching a resource
    When I send "PATCH" request to "https://jsonplaceholder.typicode.com/posts/1" with:
      | title  | body      |
      | foo222 | test body |
    Then the response code should be 200
    And the response should match json:
      """
      {
          "userId": 1,
          "id": 1,
          "title": "foo222",
          "body": "test body"
      }
      """

  @five
  Scenario: Should allow deleting a resource
    When I send "DELETE" request to "https://jsonplaceholder.typicode.com/posts/1" without body
    Then the response code should be 200

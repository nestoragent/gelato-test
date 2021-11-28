package stepDefinitions

import (
	"context"
	"encoding/json"
	"fmt"
	api "gelato-test/src/helpers"
	"github.com/cucumber/godog"
)

func InitializeTestSuite(ctx *godog.TestSuiteContext) {}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		api.ResetResponse()
		return ctx, nil
	})

	ctx.Step(`^I send "(GET|POST|PUT|DELETE|PATCH)" request to "([^"]*)" without body$`, iSendRequestToWithoutBody)
	ctx.Step(`^I send "(GET|POST|PUT|DELETE|PATCH)" request to "([^"]*)" with:$`, iSendRequestToWith)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, theResponseShouldMatchJSON)

}

func iSendRequestToWithoutBody(method, endpoint string) {
	api.SendRequestTo(method, endpoint, nil)
}

func theResponseCodeShouldBe(code int) {
	err := api.ResponseCodeShouldBe(code)
	if err != nil {
		panic(fmt.Sprintf("Response Code is not %v", code))
	}
}

func theResponseShouldMatchJSON(body *godog.DocString) {
	err := api.ResponseShouldMatchJSON(body)
	if err != nil {
		panic(fmt.Sprintf("Response is not match to json %v, err: %v", body, err))
	}
}

func iSendRequestToWith(method, endpoint string, data *godog.Table) {
	head := data.Rows[0].Cells
	values := data.Rows[1].Cells
	var requestBody = map[string]string{}

	for i := 0; i < len(head); i++ {
		requestBody[head[i].Value] = values[i].Value
	}

	bodyMarshal, _ := json.Marshal(requestBody)
	api.SendRequestTo(method, endpoint, bodyMarshal)
}

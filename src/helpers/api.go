package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"io/ioutil"
	"net/http"
)

var (
	actualResponse *http.Response
	actualBody     []byte
	contentType    = "application/json; charset=UTF-8"
)

func executeGet(endpoint string) {
	resp, err := http.Get(endpoint)
	if err != nil {
		panic(fmt.Sprintf("Error when execute GET request, error: %v", err))
	}

	actualResponse = resp
	actualBody, _ = ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
}

func executePost(endpoint string, requestBoby []byte) {
	resp, err := http.Post(endpoint, contentType, bytes.NewBuffer(requestBoby))
	//Handle Error
	if err != nil {
		panic(fmt.Sprintf("An Error Occured %v", err))
	}

	actualResponse = resp
	actualBody, _ = ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
}

func executePut(endpoint string, requestBoby []byte) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(requestBoby))
	if err != nil {
		panic(fmt.Sprintf("An Error Occured %v", err))
	}

	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("An Error Occured with Do %v", err))
	}

	actualResponse = resp
	actualBody, _ = ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
}

func executePatch(endpoint string, requestBoby []byte) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, endpoint, bytes.NewBuffer(requestBoby))
	if err != nil {
		panic(fmt.Sprintf("An Error Occured %v", err))
	}

	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("An Error Occured with Do %v", err))
	}

	actualResponse = resp
	actualBody, _ = ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
}

func executeDelete(endpoint string) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		panic(fmt.Sprintf("An Error Occured %v", err))
	}

	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("An Error Occured with Do %v", err))
	}

	actualResponse = resp
	actualBody, _ = ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
}

func ResponseCodeShouldBe(code int) error {
	if code != actualResponse.StatusCode {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, actualResponse.StatusCode)
	}
	return nil
}

func ResponseShouldMatchJSON(body *godog.DocString) error {
	var expected, actual []byte
	var data interface{}
	var actualData interface{}
	var err error

	//expected
	if err = json.Unmarshal([]byte(body.Content), &data); err != nil {
		return err
	}
	if expected, err = json.Marshal(data); err != nil {
		return err
	}

	//actual
	if err = json.Unmarshal(actualBody, &actualData); err != nil {
		return err
	}
	if actual, err = json.Marshal(actualData); err != nil {
		return err
	}

	if !bytes.Equal(actual, expected) {
		err = fmt.Errorf("expected json, does not match actual: %s", string(actualBody))
	}
	return err
}

func ResetResponse() {
	actualResponse = nil
	actualBody = nil
}

func SendRequestTo(method, endpoint string, body []byte) {
	switch method {
	case "GET":
		executeGet(endpoint)
	case "DELETE":
		executeDelete(endpoint)
	case "POST":
		executePost(endpoint, body)
	case "PUT":
		executePut(endpoint, body)
	case "PATCH":
		executePatch(endpoint, body)

	default:
		panic(fmt.Sprintf("Unknown method type: %v", method))

	}
}

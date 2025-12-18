package main

import (
	"fmt"
	"testing"
)

func TestFetchData(t *testing.T) {
	type testCase struct {
		inputUrl       string
		expectedError  error
		expectedStatus int
	}

	runCases := []testCase{
		{
			inputUrl:       "https://api.boot.dev/v1/courses_rest_api/learn-http/issues",
			expectedError:  nil,
			expectedStatus: 200,
		},
		{
			inputUrl:       "https://api.boot.dev/v1/wrong-path",
			expectedError:  fmt.Errorf("non-OK HTTP status: 404 Not Found"),
			expectedStatus: 404,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			inputUrl:       "://example.com",
			expectedError:  fmt.Errorf("network error: parse \"://example.com\": missing protocol scheme"),
			expectedStatus: 0,
		},
		{
			inputUrl:       "https://api.boot.dev/v1/courses_rest_api/learn-http/projects",
			expectedError:  nil,
			expectedStatus: 200,
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		actualStatus, actualError := fetchData(test.inputUrl)
		if (actualError == nil && test.expectedError != nil) || (actualError != nil && test.expectedError == nil) || (actualError != nil && test.expectedError != nil && actualError.Error() != test.expectedError.Error()) || actualStatus != test.expectedStatus {
			failCount++
			t.Errorf(`---------------------------------
URL:		%v
Expecting:  Error: %v - Status code: %v
Actual:     Error: %v - Status code: %v
Fail
`, test.inputUrl, test.expectedError, test.expectedStatus, actualError, actualStatus)

		} else {
			passCount++
			fmt.Printf(`---------------------------------
URL:		%v
Expecting:  Error: %v - Status code: %v
Actual:     Error: %v - Status code: %v
Pass
`, test.inputUrl, test.expectedError, test.expectedStatus, actualError, actualStatus)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}

}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

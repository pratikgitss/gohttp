package main

import (
	"fmt"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		id       string
		expected int
	}

	runCases := []testCase{
		{"0194fdc2-fa2f-4cc0-81d3-ff12045b73c8", http.StatusOK},
	}

	submitCases := append(runCases, []testCase{
		{"invalid-id", http.StatusNotFound},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users/" + test.id
		code := getUserCode(url)

		if code != test.expected {
			failCount++
			t.Errorf(`---------------------------------
ID:         %s
URL:        %s
Expecting:  %v
Actual:     %v
Fail
`, test.id, url, test.expected, code)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
ID:         %s
URL:        %s
Expecting:  %v
Actual:     %v
Pass
`, test.id, url, test.expected, code)
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

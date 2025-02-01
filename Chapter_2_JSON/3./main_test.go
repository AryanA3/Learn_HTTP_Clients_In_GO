package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		url      string
		expected []Issue
	}

	tests := []testCase{
		{
			"https://api.boot.dev/v1/courses_rest_api/learn-http/issues?limit=1",
			[]Issue{{Title: "Fix that one bug nobody understands", Estimate: 19}},
		},
		{
			"https://api.boot.dev/v1/courses_rest_api/learn-http/issues?limit=2",
			[]Issue{
				{Title: "Fix that one bug nobody understands", Estimate: 19},
				{Title: "Implement user authentication flow", Estimate: 6},
			},
		},
	}

	// Additional test cases for submission
	if withSubmit {
		tests = append(tests, []testCase{
			{"", nil},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		items, _ := getIssues(test.url)

		if !reflect.DeepEqual(items, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
URL:		%v
Expecting:  %+v
Actual:     %+v
Fail`, test.url, test.expected, items)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
URL:		%v
Expecting:  %+v
Actual:     %+v
Pass
`, test.url, test.expected, items)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

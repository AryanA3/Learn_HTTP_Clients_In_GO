package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		inputUrl string
		expected ParsedURL
	}

	tests := []testCase{
		{
			"http://waynelagner:pwn3d@jello.app:8080/boards?sort=createdAt#id",
			ParsedURL{
				protocol: "http",
				username: "waynelagner",
				password: "pwn3d",
				hostname: "jello.app",
				port:     "8080",
				pathname: "/boards",
				search:   "sort=createdAt",
				hash:     "id",
			},
		},
		{
			"https://jello.app/issues?sort=priority",
			ParsedURL{
				protocol: "https",
				username: "",
				password: "",
				hostname: "jello.app",
				port:     "",
				pathname: "/issues",
				search:   "sort=priority",
				hash:     "",
			},
		},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"", ParsedURL{}},
			{"://example.com", ParsedURL{}},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		parsedUrl := newParsedURL(test.inputUrl)
		if !reflect.DeepEqual(parsedUrl, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
URL:		%v
Expecting:  %+v
Actual:     %+v
Fail
`, test.inputUrl, test.expected, parsedUrl)

		} else {
			passCount++
			fmt.Printf(`---------------------------------
URL:		%v
Expecting:  %+v
Actual:     %+v
Pass
`, test.inputUrl, test.expected, parsedUrl)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true

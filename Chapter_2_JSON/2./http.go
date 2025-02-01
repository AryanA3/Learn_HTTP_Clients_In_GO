package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	issues := make([]Issue, 0)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&issues); err != nil {
		return nil, errors.New("")
	}

	return issues, nil
}

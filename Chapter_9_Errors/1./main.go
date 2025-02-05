package main

import (
	"errors"
	"fmt"
	"net/http"
)

func fetchData(url string) (int, error) {
	res, err := http.Get(url)

	if err != nil {
		return 0, errors.New(fmt.Sprintf("network error: %v", err.Error()))
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return res.StatusCode, errors.New(fmt.Sprintf("non-OK HTTP status: %s", res.Status))
	}

	return res.StatusCode, nil
}

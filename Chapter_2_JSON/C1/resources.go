package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}

	defer res.Body.Close()

	var data []map[string]interface{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil

}

func logResources(resources []map[string]any) {
	formattedStrings := make([]string, 0)

	for key, val := range resources {
		formattedStrings = append(formattedStrings, fmt.Sprintf("Key: %s - Value %v", key, val))
	}

	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}

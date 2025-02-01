package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	var issues []Issue
	if err := json.Unmarshal(data, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

/*

The Decode method of json.Decoder streams data from an io.Reader
into a Go struct, while json.Unmarshal works with data that's already in
[]byte format. Using a json.Decoder can be more memory-efficient because it
doesn't load all the data into memory at once.
json.Unmarshal is ideal for small JSON data you already have in memory.
When dealing with HTTP requests and responses, you will likely use
json.Decoder since it works directly with an io.Reader.

*/

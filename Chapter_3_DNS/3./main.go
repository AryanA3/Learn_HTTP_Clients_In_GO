package main

import (
	"net/url"
)

func getDomainNameFromURL(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return parsed.Hostname(), nil
}

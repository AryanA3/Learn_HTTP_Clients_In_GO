package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIPAddress(domain string) (string, error) {
	url := fmt.Sprintf("https://cloudflare-dns.com/dns-query?name=%s&type=A", domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// look at this in its file

	var dnsRes DNSResponse
	if err := json.Unmarshal(body, &dnsRes); err != nil {
		return "", fmt.Errorf("error unmarshalling json: %w", err)
	}

	if len(dnsRes.Answer) == 0 {
		return "", fmt.Errorf("no answer found")
	}

	return dnsRes.Answer[0].Data, nil
}

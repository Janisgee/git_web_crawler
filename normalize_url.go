package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedUrl, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	hostname := parsedUrl.Hostname()
	path := parsedUrl.Path

	fullParsedURL := hostname + path

	fullParsedURL = strings.ToLower(fullParsedURL)
	fullParsedURL = strings.TrimSuffix(fullParsedURL, "/")

	return fullParsedURL, nil
}

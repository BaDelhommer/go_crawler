package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	trimmed := strings.Trim(rawURL, "/")
	lower := strings.ToLower(trimmed)
	parsedURL, err := url.Parse(lower)
	if err != nil {
		return "", err
	}

	hostName := parsedURL.Hostname()
	urlPath := parsedURL.Path
	final := hostName + urlPath

	return final, nil
}

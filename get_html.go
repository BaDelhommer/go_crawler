package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHtml(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode > 399 {
		return "", errors.New(resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("wrong content type: %s", contentType)
	}

	docBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	doc := string(docBytes)

	return doc, nil
}

package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func getHtml(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode > 399 {
		return "", errors.New(resp.Status)
	}

	if val, ok := resp.Header["Content-Type"]; ok {
		if val[0] != "text/html" {
			return "", fmt.Errorf("wrong content type: %v", val[0])
		}
	}

	docBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	doc := string(docBytes)

	return doc, nil
}

package main

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	baseURL, err := url.Parse("https://boot.dev")
	if err != nil {
		fmt.Println("Couldn't parse url")
		return
	}
	tests := []struct {
		name       string
		sampleHTML string
		baseURL    *url.URL
		expected   []string
	}{
		{
			name: "normal case",
			sampleHTML: `
			<html>
        <body>
        <a href="https://boot.dev">boot.dev</a>
        <a href="/blog">Blog</a>
        <a href="https://example.com">Example</a>
        </body>
    </html>
			`,
			baseURL:  baseURL,
			expected: []string{"https://boot.dev/", "https://boot.dev/blog", "https://example.com/"},
		},
		{
			name: "no links",
			sampleHTML: `
			<html>
        <body>
        <p>No links here</p>
        </body>
    </html>
			`,
			baseURL:  baseURL,
			expected: []string{},
		},
		{
			name: "malformed html",
			sampleHTML: `
    <html>
        <body>
        <a href="https://boot.dev">boot.dev</a>
        <a href="/blog">Blog</a>
    `,
			baseURL:  baseURL,
			expected: []string{"https://boot.dev/", "https://boot.dev/blog"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsfromHTML(tc.sampleHTML, tc.baseURL)
			if err != nil {
				t.Errorf("Test %v - '%s'  FAIL: unexpected err: %v", i, tc.name, err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

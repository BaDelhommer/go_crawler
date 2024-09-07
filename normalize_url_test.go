package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	expected := "blog.boot.dev/path"
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: expected,
		},
		{
			name:     "remove slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: expected,
		},
		{
			name:     "normalize capitals",
			inputURL: "https://BLOG.boot.DEV/path",
			expected: expected,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s'  FAIL: unexpected err: %v", i, tc.name, err)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

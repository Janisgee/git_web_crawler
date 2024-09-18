package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		errorContains string
	}{{
		name:     "remove scheme",
		inputURL: "https://blog.boot.dev/path",
		expected: "blog.boot.dev/path",
	}, {
		name:     "remove trailing slash",
		inputURL: "http://blog.boot.dev/path/",
		expected: "blog.boot.dev/path",
	}, {
		name:     "lowercase capital letters",
		inputURL: "https://BLOG.boot.dev/PATH",
		expected: "blog.boot.dev/path",
	}, {
		name:     "remove scheme and trailing slash and lowercase capital letters",
		inputURL: "http://BLOG.boot.dev/path/",
		expected: "blog.boot.dev/path",
	}, {
		name:          "handle invalid URL",
		inputURL:      "://BLOG.boot.dev/path/",
		expected:      "",
		errorContains: "couldn't parse URL",
	},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %d - %s Fail: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %d - %s Fail: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %d - %s Fail: unexpected error: %v, gone none.", i, tc.name, err)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %d - %s Fail: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}

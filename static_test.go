package golem

import (
	"net/http"
	"testing"
)

func TestStatic(t *testing.T) {
	r := New()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic")
		}
	}()

	r.Static("/", http.Dir("."))
}

func TestCheckAndAppendWildcard(t *testing.T) {
	testcases := []struct {
		desc     string
		path     string
		expected string
	}{
		{
			desc:     "path ends with /",
			path:     "/",
			expected: "/*filepath",
		},
		{
			desc:     "path does not end with /",
			path:     "/somepath",
			expected: "/somepath/*filepath",
		},
	}

	for _, test := range testcases {
		gotten := checkAndAppendWildcard(test.path)

		if gotten != test.expected {
			t.Errorf("Expected %s, Gotten %s\n", test.expected, gotten)
		}
	}
}

package router

import (
	"net/http"
	"testing"
)

var (
	handler = func(r1 Request, r2 Response) {}
)

func TestAddPathToHandlers(t *testing.T) {
	testcases := []struct {
		description   string
		path          string
		handler       Handler
		expectedMaps  int
		requestedPath string
	}{
		{
			description:   "it should add a simple path",
			path:          "/user/abc",
			handler:       handler,
			requestedPath: "/user/abc",
		},
		{
			description:   "it should add a path with parameters",
			path:          "/product/:name",
			handler:       handler,
			requestedPath: "/product/abc",
		},
	}

	for _, test := range testcases {
		addPathToHandlers(http.MethodGet, test.path, test.handler)

		handler := pathMatcher(http.MethodGet, test.requestedPath)
		if handler == nil {
			t.Errorf("Handler should not be nil, handler: %v\n", handler)
		}
	}
}

func TestSanitizePath(t *testing.T) {
	testcases := []struct {
		description string
		path        string
		expected    string
	}{
		{
			description: "it should return / if path is only /",
			path:        "/",
			expected:    "/",
		},
		{
			description: "it should return the same path if there is no action required",
			path:        "/user/abc",
			expected:    "/user/abc",
		},
		{
			description: "it should remove trailing /",
			path:        "/user/abc/",
			expected:    "/user/abc",
		},
		{
			description: "it should return a sanitized path if there is one parameter",
			path:        "/user/:name",
			expected:    "/user/*",
		},
		{
			description: "it should return a sanitized path if there are multiple parameters",
			path:        "/user/:a/:b/:c",
			expected:    "/user/*/*/*",
		},
		{
			description: "it should return a sanitized path if there are multiple parameters and trailing slash",
			path:        "/user/:a/:b/:c/",
			expected:    "/user/*/*/*",
		},
	}

	for _, test := range testcases {
		gotten, _ := sanitizePath(test.path)
		if gotten != test.expected {
			t.Errorf("Expected %q, gotten %q\n", test.expected, gotten)
		}
	}
}

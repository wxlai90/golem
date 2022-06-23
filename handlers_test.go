package golem

import (
	"net/http"
	"net/url"
	"testing"
)

func TestNotFoundHandler(t *testing.T) {
	expected := http.StatusNotFound
	nf := notFoundHandler{}
	req := &http.Request{
		URL: &url.URL{
			Path: "/notfound",
		},
		Body: &mockBody{},
	}

	rw := &mockResponseWriter{}

	nf.ServeHTTP(rw, req)

	if rw.statusCode != expected {
		t.Errorf("Expected %d, Gotten %d\n", expected, rw.statusCode)
	}
}

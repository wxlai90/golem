package golem

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"testing"
)

type phone struct {
	Name string
}

type mockBody struct{}

func (m *mockBody) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

func (m *mockBody) Close() error {
	return nil
}

type mockErrorBody struct{}

func (m *mockErrorBody) Read(p []byte) (n int, err error) {
	return 0, errors.New("mock error for testing Read(), ignore.")
}

func (m *mockErrorBody) Close() error {
	return nil
}

func TestParseQueries(t *testing.T) {
	expected := "abc"
	req := &http.Request{
		URL: &url.URL{
			RawQuery: "name=abc",
		},
		Body: &mockBody{},
	}
	r := NewRequest(req)
	gotten := r.Query["name"]

	if gotten != expected {
		t.Errorf("Expected %s, Gotten %s\n", expected, gotten)
	}
}

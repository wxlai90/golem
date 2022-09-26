package golem

import (
	"errors"
	"fmt"
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

func TestUnmarshal(t *testing.T) {
	req := &http.Request{
		URL:  &url.URL{},
		Body: &mockBody{},
	}
	r := NewRequest(req)

	expected := "a"

	p := &phone{}
	r.Body = &Body{}
	r.Body.RawBytes = []byte(fmt.Sprintf(`{"name":"%s"}`, expected))
	r.Body.Unmarshal(p)

	if p.Name != "a" {
		t.Errorf("Expected %s, Gotten %s\n", expected, p.Name)
	}
}

func TestParseCookies(t *testing.T) {
	expected := "abc"
	req := &http.Request{
		URL:  &url.URL{},
		Body: &mockBody{},
		Header: map[string][]string{
			"Cookie": {fmt.Sprintf("name=%s;", expected)},
		},
	}
	r := NewRequest(req)
	gotten, _ := r.Cookies["name"]

	if gotten != expected {
		t.Errorf("Expected %s, Gotten %s\n", expected, gotten)
	}
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
	gotten, _ := r.Query["name"]

	if gotten != expected {
		t.Errorf("Expected %s, Gotten %s\n", expected, gotten)
	}
}

func TestParseRequestBody(t *testing.T) {
	req := &http.Request{
		URL:  &url.URL{},
		Body: &mockErrorBody{},
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Req body read should panic but did not.")
		}
	}()
	NewRequest(req)
}

func TestBag(t *testing.T) {
	expected := "abc"
	req := &http.Request{
		URL:  &url.URL{},
		Body: &mockBody{},
	}
	r := NewRequest(req)
	r.Put("name", expected)

	if gotten, ok := r.Get("name"); ok {
		if gotten != expected {
			t.Errorf("Expected %s, Gotten %s\n", expected, gotten)
		}
	}

	_, ok := r.Get("not found")
	if ok {
		t.Errorf("Expected false, Gotten true\n")
	}
}

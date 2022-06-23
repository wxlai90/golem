package golem

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/julienschmidt/httprouter"
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
	params := httprouter.Params{}
	r := NewRequest(req, params)

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
	params := httprouter.Params{}
	r := NewRequest(req, params)
	gotten, _ := r.Cookies["name"]

	if gotten != expected {
		t.Errorf("Expected %s, Gotten %s\n", expected, gotten)
	}
}

func TestParseParams(t *testing.T) {
	expected := "abc"
	req := &http.Request{
		URL:  &url.URL{},
		Body: &mockBody{},
	}
	params := httprouter.Params{
		{
			Key:   "name",
			Value: expected,
		},
	}
	r := NewRequest(req, params)
	gotten, _ := r.Params["name"]

	if gotten != expected {
		t.Errorf("Expected %s, Gotten %s\n", expected, gotten)
	}
}

func TestParseQueries(t *testing.T) {
	expected := "abc"
	req := &http.Request{
		URL: &url.URL{
			RawQuery: "name=abc;",
		},
		Body: &mockBody{},
	}
	params := httprouter.Params{}
	r := NewRequest(req, params)
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
	params := httprouter.Params{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Req body read should panic but did not.")
		}
	}()
	NewRequest(req, params)
}

func TestBag(t *testing.T) {
	expected := "abc"
	req := &http.Request{
		URL:  &url.URL{},
		Body: &mockBody{},
	}
	params := httprouter.Params{}
	r := NewRequest(req, params)
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

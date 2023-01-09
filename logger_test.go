package golem

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	httpReq := &http.Request{
		URL: &url.URL{
			Path:     "/user",
			RawQuery: "",
		},
		Header: map[string][]string{},
		Method: http.MethodGet,
		Body:   &mockBody{},
	}

	m := &mockResponseWriter{}
	req := NewRequest(httpReq)
	res := NewResponse(m)

	var isNextCalled bool
	DefaultLogger()(req, res, func() {
		isNextCalled = true
	})

	if !isNextCalled {
		t.Errorf("Expected next() to be called regardless.\n")
	}

	gotten := buf.String()[20:] // trim timestamp
	expected := "[GET] - /user\n"
	if expected != gotten {
		t.Errorf("expected %s, gotten %s\n", expected, gotten)
	}

}

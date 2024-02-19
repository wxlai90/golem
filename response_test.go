package golem

import (
	"net/http"
	"testing"
)

type mockResponseWriter struct {
	written    []byte
	header     map[string][]string
	statusCode int
}

func (m *mockResponseWriter) Header() http.Header {
	if m.header == nil {
		m.header = map[string][]string{}
	}

	return m.header
}

func (m *mockResponseWriter) Write(data []byte) (int, error) {
	m.written = data
	return 0, nil
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.statusCode = statusCode
}

func TestNewResponse(t *testing.T) {
	r := NewResponse(nil)
	if r == nil {
		t.Errorf("NewResponse is nil")
	}
}

func TestJson(t *testing.T) {
	expected := []byte(`{"Name":"abc"}`)
	expected = append(expected, 0xa)
	m := &mockResponseWriter{}
	r := NewResponse(m)
	payload := struct {
		Name string
	}{
		Name: "abc",
	}
	r.Json(payload)

	if string(m.written) != string(expected) {
		t.Errorf("Expected %s, Gotten %s\n", expected, m.written)
	}

	if contentType, ok := m.header["Content-Type"]; ok {
		if contentType[0] != "application/json" {
			t.Errorf("Expected content-type to be application/json")
		}
	}
}

func TestSend(t *testing.T) {
	expected := "abc"
	m := &mockResponseWriter{}
	r := NewResponse(m)
	r.Send("abc")

	if string(m.written) != expected {
		t.Errorf("Expected %s, Gotten %s\n", expected, m.written)
	}
}

func TestCookie(t *testing.T) {
	expected := "name=abc"
	m := &mockResponseWriter{}
	r := NewResponse(m)
	r.Cookie("name", "abc")

	if cookie, ok := m.header["Set-Cookie"]; ok {
		if cookie[0] != expected {
			t.Errorf("Expected %s, Gotten %s\n", expected, cookie[0])
		}
	}
}

func TestStatus(t *testing.T) {
	expected := http.StatusOK
	m := &mockResponseWriter{}
	r := NewResponse(m)
	r.Status(expected).Send("")

	if m.statusCode != expected {
		t.Errorf("Expected %d, Gotten %d\n", expected, m.statusCode)
	}
}

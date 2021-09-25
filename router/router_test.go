package router

import (
	"io"
	"net/http"
	"testing"
)

type MockWriter struct {
	response string
}

func (m *MockWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *MockWriter) Write(p []byte)(int, error) {
	m.response = string(p)
	return 1, nil
}

func (m *MockWriter) WriteString(s string) {
}

func (m *MockWriter) WriteHeader(int) {}


type expectedResponse struct {
	method string
	response string
}


var testMethods []expectedResponse = []expectedResponse{
	{http.MethodGet, "GET okay!"},
	{http.MethodPost, "POST okay!"},
}

func TestHTTPMethods(t *testing.T){
	r := New()



	r.GET("/", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "GET okay!")
	})

	r.POST("/", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "POST okay!")
	})

	for _, test := range testMethods {
		w := new(MockWriter)

		req, _ := http.NewRequest(test.method, "/", nil)

		r.ServeHTTP(w, req)

		if test.response != w.response {
			t.Errorf("Expected %s, Gotten %s\n", test.response, w.response)
		}
	}
}

func TestNotFound(t *testing.T){
	r := New()
	w := new(MockWriter)

	test := expectedResponse{
		http.MethodGet,
		"404 page not found\n",
	}

	req, _ := http.NewRequest(http.MethodGet, "/not/found",  nil)

	r.ServeHTTP(w, req)

	if test.response != w.response {
		t.Errorf("Expected %s, Gotten %s\n", test.response, w.response)
	}
}
package golem

import (
	"net/http"
	"net/url"
	"testing"
)

func TestRoutes(t *testing.T) {
	testcases := []struct {
		desc        string
		shouldPanic bool
		prefix      string
		path        string
	}{
		{
			desc:        "Should not panic when prefix and paths are correct",
			shouldPanic: false,
			prefix:      "/prefix",
			path:        "/path",
		},
		{
			desc:        "Should panic when prefix is incorrect",
			shouldPanic: true,
			prefix:      "abc",
			path:        "def",
		},
	}

	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodHead,
		http.MethodOptions,
	}

	for _, test := range testcases {
		for _, method := range methods {
			r := New()
			sub := &SubRouter{}
			sub.nodes = []subRouterNode{
				{
					path:        test.path,
					method:      method,
					handler:     func(req *Request, res *Response) {},
					middlewares: []Middleware{},
				},
			}

			defer func() {
				if r := recover(); r != nil {
					if !test.shouldPanic {
						t.Errorf(test.desc)
					}
				}

				if r == nil {
					if test.shouldPanic {
						t.Errorf(test.desc)
					}
				}
			}()

			r.Routes(test.prefix, sub)
		}
	}
}

// func TestAdapter(t *testing.T) {
// 	expected := "abc"
// 	handlerCalled := false
// 	middlewareCalled := false
// 	wrapped := adapter(func(req *Request, res *Response) {
// 		handlerCalled = true
// 		res.Send(expected)
// 	}, []middleware{
// 		func(req *Request, res *Response, next Next) {
// 			next()
// 		},
// 		func(req *Request, res *Response, next Next) {
// 			middlewareCalled = true
// 			next()
// 		},
// 	})

// 	if wrapped == nil {
// 		t.Errorf("Expected wrapped handler from adapter, Gotten nil.\n")
// 	}

// 	r := &http.Request{
// 		URL:  &url.URL{},
// 		Body: &mockBody{},
// 	}
// 	rw := &mockResponseWriter{}
// 	adapterRunning := make(chan struct{})
// 	adapterDone := make(chan struct{})
// 	go func() {
// 		close(adapterRunning)
// 		wrapped(rw, r, nil)
// 		defer close(adapterDone)
// 	}()

// 	<-adapterRunning
// 	if !handlerCalled {
// 		t.Errorf("Expected handler to be called.")
// 	}

// 	if !middlewareCalled {
// 		t.Errorf("Expected middlewares to be called.")
// 	}
// 	<-adapterDone
// }

func TestServeHTTP(t *testing.T) {
	ro := New()
	r := &http.Request{
		URL:  &url.URL{},
		Body: &mockBody{},
	}
	rw := &mockResponseWriter{}

	ro.ServeHTTP(rw, r)

	if rw.statusCode != http.StatusNotFound {
		t.Errorf("Expected %d, Gotten %d\n", http.StatusNotFound, rw.statusCode)
	}
}

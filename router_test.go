package golem

import (
	"net/http"
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

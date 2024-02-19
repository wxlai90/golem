package golem

import (
	"testing"
)

var (
	router         *Router
	mockMiddleware = func(req *Request, res *Response, next Next) {}
	contMiddleware = func(req *Request, res *Response, next Next) { next() }
	stopMiddleware = func(req *Request, res *Response, next Next) {}
)

func TestUse(t *testing.T) {
	testcases := []struct {
		desc        string
		middlewares []Middleware
	}{
		{
			desc: "should register middleware when middlewares are defined",
			middlewares: []Middleware{
				mockMiddleware,
			},
		},
		{
			desc: "should add more nodes into middlewares slice",
			middlewares: []Middleware{
				mockMiddleware,
				mockMiddleware,
			},
		},
	}

	for _, test := range testcases {
		t.Run(test.desc, func(t *testing.T) {
			globalMiddlewares = []Middleware{}
			for _, middleware := range test.middlewares {
				router.Use(middleware)
			}

			if len(globalMiddlewares) != len(test.middlewares) {
				t.Errorf("expected %d, gotten %d\n", len(test.middlewares), len(globalMiddlewares))
			}
		})
	}
}

func TestTraverseGlobalMiddlewares(t *testing.T) {
	req := &Request{}
	res := &Response{}

	testcases := []struct {
		desc           string
		middleware     Middleware
		shouldContinue bool
	}{
		{
			desc:           "should call middlewares fully if next is called",
			middleware:     contMiddleware,
			shouldContinue: true,
		},
		{
			desc:           "should call middlewares but stop if next is not called",
			middleware:     stopMiddleware,
			shouldContinue: false,
		},
	}

	for _, test := range testcases {
		t.Run(test.desc, func(t *testing.T) {
			globalMiddlewares = []Middleware{
				test.middleware,
			}

			gotten := traverseGlobalMiddlewares(req, res)

			if gotten != test.shouldContinue {
				t.Errorf("Expected %v, Gotten %v\n", test.shouldContinue, gotten)
			}
		})
	}
}

func TestTraverseLocalMiddlewares(t *testing.T) {
	req := &Request{}
	res := &Response{}

	contMiddlewares := []Middleware{
		contMiddleware,
		contMiddleware,
		contMiddleware,
	}

	stopMiddlewares := []Middleware{
		contMiddleware,
		stopMiddleware,
		contMiddleware,
	}

	testcases := []struct {
		desc           string
		middlewares    []Middleware
		shouldContinue bool
	}{
		{
			desc:           "should call middlewares fully if next is called",
			middlewares:    contMiddlewares,
			shouldContinue: true,
		},
		{
			desc:           "should call middlewares but stop if next is not called",
			middlewares:    stopMiddlewares,
			shouldContinue: false,
		},
	}

	for _, test := range testcases {
		t.Run(test.desc, func(t *testing.T) {
			gotten := traverseMiddlewares(req, res, test.middlewares)

			if gotten != test.shouldContinue {
				t.Errorf("Expected %v, Gotten %v\n", test.shouldContinue, gotten)
			}
		})
	}
}

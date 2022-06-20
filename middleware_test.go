package golem

import "testing"

var (
	router         *Router
	mockMiddleware = func(req *Request, res *Response, next Next) {}
	contMiddleware = func(req *Request, res *Response, next Next) { next() }
	stopMiddleware = func(req *Request, res *Response, next Next) {}
)

type mock struct{}

func TestUse(t *testing.T) {
	testcases := []struct {
		desc        string
		middlewares []middleware
	}{
		{
			desc: "should initialize linkedlist when there are middlewares",
			middlewares: []middleware{
				mockMiddleware,
			},
		},
		{
			desc: "should add more nodes into middlewares linkedlist",
			middlewares: []middleware{
				mockMiddleware,
				mockMiddleware,
			},
		},
	}

	for _, test := range testcases {
		t.Run(test.desc, func(t *testing.T) {
			for _, middleware := range test.middlewares {
				router.Use(middleware)
			}

			if head == nil || curr == nil {
				t.Errorf("middlewares uninitialized")
			}
		})
	}
}

func TestTraverseGlobalMiddlewares(t *testing.T) {
	req := &Request{}
	res := &Response{}

	contLL := &middlewareNode{
		middleware: contMiddleware,
	}

	contLL.next = &middlewareNode{
		middleware: contMiddleware,
	}

	contLL.next.next = &middlewareNode{
		middleware: contMiddleware,
	}

	stopLL := &middlewareNode{
		middleware: contMiddleware,
	}

	stopLL.next = &middlewareNode{
		middleware: stopMiddleware,
	}

	stopLL.next.next = &middlewareNode{
		middleware: contMiddleware,
	}

	testcases := []struct {
		desc           string
		head           *middlewareNode
		shouldContinue bool
	}{
		{
			desc:           "should call middlewares fully if next is called",
			head:           contLL,
			shouldContinue: true,
		},
		{
			desc:           "should call middlewares but stop if next is not called",
			head:           stopLL,
			shouldContinue: false,
		},
	}

	for _, test := range testcases {
		t.Run(test.desc, func(t *testing.T) {
			head = test.head
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

	contMiddlewares := []middleware{
		contMiddleware,
		contMiddleware,
		contMiddleware,
	}

	stopMiddlewares := []middleware{
		contMiddleware,
		stopMiddleware,
		contMiddleware,
	}

	testcases := []struct {
		desc           string
		middlewares    []middleware
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
			gotten := traverseLocalMiddlewares(req, res, test.middlewares)

			if gotten != test.shouldContinue {
				t.Errorf("Expected %v, Gotten %v\n", test.shouldContinue, gotten)
			}
		})
	}
}

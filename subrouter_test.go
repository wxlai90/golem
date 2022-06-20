package golem

import (
	"net/http"
	"testing"
)

func TestSubRouter(t *testing.T) {
	subRouter := NewSubRouter()

	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodHead,
		http.MethodOptions,
	}

	subRouter.GET("/", func(req *Request, res *Response) {}, func(req *Request, res *Response, next Next) {})
	subRouter.POST("/", func(req *Request, res *Response) {}, func(req *Request, res *Response, next Next) {})
	subRouter.PUT("/", func(req *Request, res *Response) {}, func(req *Request, res *Response, next Next) {})
	subRouter.PATCH("/", func(req *Request, res *Response) {}, func(req *Request, res *Response, next Next) {})
	subRouter.DELETE("/", func(req *Request, res *Response) {}, func(req *Request, res *Response, next Next) {})
	subRouter.HEAD("/", func(req *Request, res *Response) {}, func(req *Request, res *Response, next Next) {})
	subRouter.OPTIONS("/", func(req *Request, res *Response) {}, func(req *Request, res *Response, next Next) {})

	for i, node := range subRouter.nodes {
		if node.method != methods[i] {
			t.Errorf("Expected %s, Gotten %s\n", methods[i], node.method)
		}
	}
}

func TestSubRouterUse(t *testing.T) {
	expected := 2
	subRouter := NewSubRouter()
	for i := 0; i < expected; i++ {
		subRouter.Use(func(req *Request, res *Response, next Next) {})
	}

	if len(subRouter.globalMiddlewares) != expected {
		t.Errorf("Expected %d, Gotten %d\n", expected, len(subRouter.globalMiddlewares))
	}
}

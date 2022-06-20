package golem

import "net/http"

type subRouterNode struct {
	path        string
	method      string
	handler     Handler
	middlewares []middleware
}

type SubRouter struct {
	globalMiddlewares []middleware
	nodes             []subRouterNode
}

func (s *SubRouter) GET(path string, handler Handler, middlewares ...middleware) {
	newNode := subRouterNode{
		path:        path,
		method:      http.MethodGet,
		handler:     handler,
		middlewares: middlewares,
	}

	s.nodes = append(s.nodes, newNode)
}

func (s *SubRouter) POST(path string, handler Handler, middlewares ...middleware) {
	newNode := subRouterNode{
		path:        path,
		method:      http.MethodPost,
		handler:     handler,
		middlewares: middlewares,
	}

	s.nodes = append(s.nodes, newNode)
}

func (s *SubRouter) PUT(path string, handler Handler, middlewares ...middleware) {
	newNode := subRouterNode{
		path:        path,
		method:      http.MethodPut,
		handler:     handler,
		middlewares: middlewares,
	}

	s.nodes = append(s.nodes, newNode)
}

func (s *SubRouter) PATCH(path string, handler Handler, middlewares ...middleware) {
	newNode := subRouterNode{
		path:        path,
		method:      http.MethodPatch,
		handler:     handler,
		middlewares: middlewares,
	}

	s.nodes = append(s.nodes, newNode)
}

func (s *SubRouter) DELETE(path string, handler Handler, middlewares ...middleware) {
	newNode := subRouterNode{
		path:        path,
		method:      http.MethodDelete,
		handler:     handler,
		middlewares: middlewares,
	}

	s.nodes = append(s.nodes, newNode)
}

func (s *SubRouter) HEAD(path string, handler Handler, middlewares ...middleware) {
	newNode := subRouterNode{
		path:        path,
		method:      http.MethodHead,
		handler:     handler,
		middlewares: middlewares,
	}

	s.nodes = append(s.nodes, newNode)
}

func (s *SubRouter) OPTIONS(path string, handler Handler, middlewares ...middleware) {
	newNode := subRouterNode{
		path:        path,
		method:      http.MethodOptions,
		handler:     handler,
		middlewares: middlewares,
	}

	s.nodes = append(s.nodes, newNode)
}

func (s *SubRouter) Use(mw middleware) {
	s.globalMiddlewares = append(s.globalMiddlewares, mw)
}

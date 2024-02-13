package golem

import "net/http"

func New() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func NewSubRouter() *SubRouter {
	return &SubRouter{
		nodes: []subRouterNode{},
	}
}

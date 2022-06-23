package golem

import (
	"github.com/julienschmidt/httprouter"
)

func New() *Router {
	r := httprouter.New()
	nf := &notFoundHandler{}
	r.NotFound = nf

	return &Router{
		InnerRouter: r,
	}
}

func NewSubRouter() *SubRouter {
	return &SubRouter{
		nodes: []subRouterNode{},
	}
}

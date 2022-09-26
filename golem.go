package golem

func New() *Router {
	return &Router{
		handlers: make(map[string]map[string]handlerNode),
	}
}

func NewSubRouter() *SubRouter {
	return &SubRouter{
		nodes: []subRouterNode{},
	}
}

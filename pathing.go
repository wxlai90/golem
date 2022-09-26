package golem

type handlerNode struct {
	handler     Handler
	middlewares []Middleware
}

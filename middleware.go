package golem

var (
	head *middlewareNode
	curr *middlewareNode
)

type middlewareNode struct {
	middleware Middleware
	next       *middlewareNode
}

type Next func()

type Middleware func(req *Request, res *Response, next Next)

// Use() exposes a middleware system
func (r *Router) Use(mw Middleware) {
	if head == nil {
		head = &middlewareNode{
			middleware: mw,
		}
		curr = head
		return
	}

	curr.next = &middlewareNode{
		middleware: mw,
	}

	curr = curr.next
}

func traverseGlobalMiddlewares(req *Request, res *Response) bool {
	start := head
	for start != nil {
		cont := false

		start.middleware(req, res, func() { cont = true })
		if !cont {
			return false
		}

		start = start.next
	}

	return true
}

func traverseLocalMiddlewares(req *Request, res *Response, middlewares []Middleware) bool {
	for _, mw := range middlewares {
		cont := false
		mw(req, res, func() {
			cont = true
		})
		if !cont {
			return false
		}
	}

	return true
}

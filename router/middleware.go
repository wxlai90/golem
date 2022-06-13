package router

var (
	head *middlewareNode
	curr *middlewareNode
)

type middlewareNode struct {
	middleware middleware
	next       *middlewareNode
}

type Next func()

type middleware func(next Next, req *Request, res *Response)

// Use() exposes a middleware system
func (r *Router) Use(mw middleware) {
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

func traverseMiddlewares(req *Request, res *Response) bool {
	start := head
	for start != nil {
		cont := false
		start.middleware(func() {
			cont = true
		}, req, res)

		if !cont {
			return false
		}

		start = start.next
	}

	return true
}

package golem

var (
	globalMiddlewares = []Middleware{}
)

type Next func()

type Middleware func(req *Request, res *Response, next Next)

func (r *Router) Use(mw Middleware) {
	globalMiddlewares = append(globalMiddlewares, mw)
}

func traverseGlobalMiddlewares(req *Request, res *Response) bool {
	return traverseMiddlewares(req, res, globalMiddlewares)
}

func traverseMiddlewares(req *Request, res *Response, middlewares []Middleware) bool {
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

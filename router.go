package golem

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	InnerRouter *httprouter.Router
}

func (r *Router) Routes(prefix string, subRouter *SubRouter) {
	for _, subRoute := range subRouter.nodes {
		subRoute.middlewares = append(subRoute.middlewares, subRouter.globalMiddlewares...)
		fullPath := fmt.Sprintf("%s%s", prefix, subRoute.path)

		switch subRoute.method {
		case http.MethodGet:
			r.GET(fullPath, subRoute.handler, subRoute.middlewares...)
		case http.MethodPost:
			r.POST(fullPath, subRoute.handler, subRoute.middlewares...)
		case http.MethodPut:
			r.PUT(fullPath, subRoute.handler, subRoute.middlewares...)
		case http.MethodPatch:
			r.PATCH(fullPath, subRoute.handler, subRoute.middlewares...)
		case http.MethodDelete:
			r.DELETE(fullPath, subRoute.handler, subRoute.middlewares...)
		case http.MethodHead:
			r.HEAD(fullPath, subRoute.handler, subRoute.middlewares...)
		case http.MethodOptions:
			r.OPTIONS(fullPath, subRoute.handler, subRoute.middlewares...)
		}
	}
}

func (r *Router) Listen(addr string, fn ...func()) {
	if len(fn) > 0 && fn[0] != nil {
		fn[0]()
	}

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	server.ListenAndServe()
}

type Handler func(req *Request, res *Response)

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ro.InnerRouter.ServeHTTP(w, r)
}

func adapter(handler Handler, middlewares []Middleware) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		req := NewRequest(r, p)
		res := NewResponse(rw)

		cont := traverseGlobalMiddlewares(req, res)
		if !cont {
			return
		}

		if len(middlewares) > 0 {
			cont = traverseLocalMiddlewares(req, res, middlewares)
			if !cont {
				return
			}
		}

		handler(req, res)
	}
}

func (r *Router) GET(path string, handler Handler, middlewares ...Middleware) {
	r.InnerRouter.GET(path, adapter(handler, middlewares))
}

func (r *Router) POST(path string, handler Handler, middlewares ...Middleware) {
	r.InnerRouter.POST(path, adapter(handler, middlewares))
}

func (r *Router) PUT(path string, handler Handler, middlewares ...Middleware) {
	r.InnerRouter.PUT(path, adapter(handler, middlewares))
}

func (r *Router) PATCH(path string, handler Handler, middlewares ...Middleware) {
	r.InnerRouter.PATCH(path, adapter(handler, middlewares))
}

func (r *Router) DELETE(path string, handler Handler, middlewares ...Middleware) {
	r.InnerRouter.DELETE(path, adapter(handler, middlewares))
}

func (r *Router) HEAD(path string, handler Handler, middlewares ...Middleware) {
	r.InnerRouter.HEAD(path, adapter(handler, middlewares))
}

func (r *Router) OPTIONS(path string, handler Handler, middlewares ...Middleware) {
	r.InnerRouter.OPTIONS(path, adapter(handler, middlewares))
}

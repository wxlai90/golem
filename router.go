package golem

import (
	"fmt"
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func (r *Router) Routes(prefix string, subRouter *SubRouter) {
	for _, subRoute := range subRouter.nodes {
		subRoute.middlewares = append(subRoute.middlewares, subRouter.globalMiddlewares...)
		fullPath := fmt.Sprintf("%s%s", prefix, subRoute.path)

		switch subRoute.method {
		case http.MethodGet:
			r.GET(fullPath, subRoute.handler)
		case http.MethodPost:
			r.POST(fullPath, subRoute.handler)
		case http.MethodPut:
			r.PUT(fullPath, subRoute.handler)
		case http.MethodPatch:
			r.PATCH(fullPath, subRoute.handler)
		case http.MethodDelete:
			r.DELETE(fullPath, subRoute.handler)
		case http.MethodHead:
			r.HEAD(fullPath, subRoute.handler)
		case http.MethodOptions:
			r.OPTIONS(fullPath, subRoute.handler)
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
type HandleFunc func(rw http.ResponseWriter, r *http.Request)

func (ro *Router) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ro.mux.ServeHTTP(rw, r)
}

func adapter(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := NewRequest(r)
		res := NewResponse(w)

		handler(req, res)
	}
}

func (r *Router) register(method, path string, handler Handler) {
	r.mux.HandleFunc(fmt.Sprintf("%s %s", method, path), adapter(handler))
}

func (r *Router) GET(path string, handler Handler) {
	r.register(http.MethodGet, path, handler)
}

func (r *Router) POST(path string, handler Handler) {
	r.register(http.MethodPost, path, handler)
}

func (r *Router) PUT(path string, handler Handler) {
	r.register(http.MethodPut, path, handler)
}

func (r *Router) PATCH(path string, handler Handler) {
	r.register(http.MethodPatch, path, handler)
}

func (r *Router) DELETE(path string, handler Handler) {
	r.register(http.MethodDelete, path, handler)
}

func (r *Router) HEAD(path string, handler Handler) {
	r.register(http.MethodHead, path, handler)
}

func (r *Router) OPTIONS(path string, handler Handler) {
	r.register(http.MethodOptions, path, handler)
}

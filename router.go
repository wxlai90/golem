package golem

import (
	"fmt"
	"net/http"
)

type Router struct {
	handlers         map[string]map[string]handlerNode
	staticFileServer http.Handler
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
type HandleFunc func(rw http.ResponseWriter, r *http.Request)

func (ro *Router) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if node, ok := ro.handlers[r.Method][r.URL.Path]; ok {
		req := NewRequest(r)
		res := NewResponse(rw)
		cont := traverseGlobalMiddlewares(req, res)
		if !cont {
			return
		}

		if len(node.middlewares) > 0 {
			cont = traverseLocalMiddlewares(req, res, node.middlewares)
			if !cont {
				return
			}
		}

		node.handler(req, res)
		return
	}

	// TODO: allow multiple static, use proper matching
	if ro.staticFileServer != nil {
		ro.staticFileServer.ServeHTTP(rw, r)
		return
	}

	http.NotFound(rw, r)
}

func (r *Router) register(method, path string, handler Handler, middlewares []Middleware) {
	if _, exists := r.handlers[method]; !exists {
		r.handlers[method] = make(map[string]handlerNode)
	}

	node := handlerNode{
		handler:     handler,
		middlewares: middlewares,
	}

	r.handlers[method][path] = node
}

func (r *Router) GET(path string, handler Handler, middlewares ...Middleware) {
	r.register(http.MethodGet, path, handler, middlewares)
}

func (r *Router) POST(path string, handler Handler, middlewares ...Middleware) {
	r.register(http.MethodPost, path, handler, middlewares)
}

func (r *Router) PUT(path string, handler Handler, middlewares ...Middleware) {
	r.register(http.MethodPut, path, handler, middlewares)
}

func (r *Router) PATCH(path string, handler Handler, middlewares ...Middleware) {
	r.register(http.MethodPatch, path, handler, middlewares)
}

func (r *Router) DELETE(path string, handler Handler, middlewares ...Middleware) {
	r.register(http.MethodDelete, path, handler, middlewares)
}

func (r *Router) HEAD(path string, handler Handler, middlewares ...Middleware) {
	r.register(http.MethodHead, path, handler, middlewares)
}

func (r *Router) OPTIONS(path string, handler Handler, middlewares ...Middleware) {
	r.register(http.MethodOptions, path, handler, middlewares)
}

// TODO: allow multiple static
func (r *Router) serveFiles(path string, dir http.Dir) {
	fileServer := http.FileServer(dir)
	r.staticFileServer = fileServer
}

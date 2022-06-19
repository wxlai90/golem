package golem

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	InnerRouter *httprouter.Router
}

func (r *Router) Listen(addr string, fn ...func()) {
	if len(fn) > 0 && fn[0] != nil {
		fn[0]()
	}

	if !strings.Contains(addr, ":") {
		// assuming only port number was passed in, append ":" to port
		addr = ":" + addr
	}

	http.ListenAndServe(addr, r)
}

type Handler func(*Request, *Response)

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ro.InnerRouter.ServeHTTP(w, r)
}

func adapter(handler Handler, middlewares []middleware) httprouter.Handle {
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

func (r *Router) GET(path string, handler Handler, middlewares ...middleware) {
	r.InnerRouter.GET(path, adapter(handler, middlewares))
}

func (r *Router) POST(path string, handler Handler, middlewares ...middleware) {
	r.InnerRouter.POST(path, adapter(handler, middlewares))
}

func (r *Router) PUT(path string, handler Handler, middlewares ...middleware) {
	r.InnerRouter.PUT(path, adapter(handler, middlewares))
}

func (r *Router) PATCH(path string, handler Handler, middlewares ...middleware) {
	r.InnerRouter.PATCH(path, adapter(handler, middlewares))
}

func (r *Router) DELETE(path string, handler Handler, middlewares ...middleware) {
	r.InnerRouter.DELETE(path, adapter(handler, middlewares))
}

func (r *Router) HEAD(path string, handler Handler, middlewares ...middleware) {
	r.InnerRouter.HEAD(path, adapter(handler, middlewares))
}

func (r *Router) OPTIONS(path string, handler Handler, middlewares ...middleware) {
	r.InnerRouter.OPTIONS(path, adapter(handler, middlewares))
}

package router

import (
	"net/http"
	"strings"
)

var handlers map[string]map[string]Node = map[string]map[string]Node{}

type Router struct{}

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

type TokenizedPath struct {
	originalPath string
	wildcardPath string
	wildcards    []string
	handler      Handler
}

type Handler func(Request, Response)

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := pathMatcher(r.Method, r.URL.Path)
	if handler == nil {
		http.NotFound(w, r)
		return
	}

	req := Request{}
	res := Response{
		W: w,
	}
	handler(req, res)
}

func (r *Router) GET(path string, handler Handler) {
	addPathToHandlers(http.MethodGet, path, handler)
}

func (r *Router) POST(path string, handler Handler) {
	addPathToHandlers(http.MethodPost, path, handler)
}

func (r *Router) PUT(path string, handler Handler) {
	addPathToHandlers(http.MethodPut, path, handler)
}

func (r *Router) PATCH(path string, handler Handler) {
	addPathToHandlers(http.MethodPatch, path, handler)
}

func (r *Router) DELETE(path string, handler Handler) {
	addPathToHandlers(http.MethodDelete, path, handler)
}

package router

import (
	"net/http"
)

var handlers map[string]map[string]Handler = map[string]map[string]Handler{}

type Router struct{}

type Handler func(http.ResponseWriter, *http.Request)

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := handlers[r.Method][r.URL.Path]; ok {
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func addPathToHandlers(method string, path string, handler Handler) {
	_, exists := handlers[method]

	if !exists {
		handlers[method] = make(map[string]Handler)
		handlers[method][path] = handler
	}
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

func New() *Router {
	return &Router{}
}

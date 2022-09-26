package golem

import (
	"net/http"
)

type notFoundHandler struct{}

func (n *notFoundHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	req := NewRequest(r)
	res := NewResponse(rw)

	cont := traverseGlobalMiddlewares(req, res)
	if !cont {
		return
	}

	res.Status(404).Send("404 Not Found")
}

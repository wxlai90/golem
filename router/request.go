package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Request struct {
	R       *http.Request
	Cookies map[string]string
	Params  map[string]string
}

func NewRequest(r *http.Request, p httprouter.Params) Request {
	req := Request{
		R: r,
	}

	req.parseCookies()
	req.parseParams(p)

	return req
}

func (r *Request) parseCookies() {
	r.Cookies = map[string]string{}
	for _, cookie := range r.R.Cookies() {
		r.Cookies[cookie.Name] = cookie.Value
	}
}

func (r *Request) parseParams(params httprouter.Params) {
	r.Params = map[string]string{}
	for _, param := range params {
		r.Params[param.Key] = param.Value
	}
}

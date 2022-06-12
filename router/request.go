package router

import (
	"net/http"
)

type Request struct {
	R       *http.Request
	Cookies map[string]string
}

func NewRequest(r *http.Request) Request {
	req := Request{
		R: r,
	}

	req.parseCookies()

	return req
}

func (r *Request) parseCookies() {
	r.Cookies = map[string]string{}
	for _, cookie := range r.R.Cookies() {
		r.Cookies[cookie.Name] = cookie.Value
	}
}

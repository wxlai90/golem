package golem

import "log"

func DefaultLogger() Middleware {
	return func(req *Request, res *Response, next Next) {
		log.Printf("[%s] - %s\n", req.Method, req.URL.Path)
		next()
	}
}

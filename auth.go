package golem

import "net/http"

func BasicAuth(username, password string) Middleware {
	return func(req *Request, res *Response, next Next) {
		uname, passwd, ok := req.BasicAuth()

		if ok {
			if username == uname && password == passwd {
				next()
				return
			}
		}

		res.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
	}
}

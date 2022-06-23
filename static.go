package golem

import "net/http"

func (r *Router) Static(path string, dir http.Dir) {
	if path[len(path)-1] != '/' {
		path += "/"
	}
	path += "*filepath" // compatibility with httprouter

	r.InnerRouter.ServeFiles(path, dir)
}

package golem

import "net/http"

func (r *Router) Static(path string, dir http.Dir) {
	http.Handle(path, http.FileServer(dir))
	r.ServeFiles(path, dir)
}

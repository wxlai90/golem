package golem

import "net/http"

func checkAndAppendWildcard(path string) string {
	if path[len(path)-1] != '/' {
		path += "/"
	}
	path += "*filepath" // compatibility with httprouter

	return path
}

func (r *Router) Static(path string, dir http.Dir) {
	path = checkAndAppendWildcard(path)
	// r.ServeFiles(path, dir)
}

package router

import (
	"fmt"
	"strings"
)

type Node struct {
	handler   Handler
	variables map[int]string
}

// pathMatcher finds a handler based on a requested url path
// with or without wildcards/params
// very bad implementation for now... doesn't even work
func pathMatcher(method, path string) Handler {
	methodHandlers := handlers[method]
	if node, ok := methodHandlers[path]; ok {
		return node.handler
	}

	segments := strings.Split(path, "/")[1:]

	for i := len(segments) - 1; i > 0; i-- {
		s := "/" + strings.Join(segments[:i], "/") + "/*"

		if node, ok := methodHandlers[s]; ok {
			fmt.Printf("Found parameter: %s to be %s\n", node.variables[i], segments[i])
			return node.handler
		}
	}

	return nil
}

func sanitizePath(path string) (string, map[int]string) {
	if path == "/" {
		return path, map[int]string{}
	}

	sanePath := ""
	i := 0
	varc := -1
	vars := map[int]string{}

	for i < len(path) {
		if path[i] == '/' {
			varc++
		}

		if path[i] == ':' {
			// means its a parameter
			sanePath += "*"
			name := ""
			i++
			for i < len(path) && path[i] != '/' {
				name += string(path[i])
				i++
			}
			vars[varc] = name
			varc++
		} else {
			sanePath += string(path[i])
			i++
		}
	}

	if sanePath[len(sanePath)-1] == '/' {
		sanePath = sanePath[:len(sanePath)-1]
	}

	return sanePath, vars
}

// addPathToHandlers adds a Handler with a path based on segments for easy retrieval
// if wildcard path exists, cannot add a duplicate concrete path
// e.g. /user/:name means /user/abc cannot be added
func addPathToHandlers(method string, path string, handler Handler) {
	_, exists := handlers[method]

	if !exists {
		handlers[method] = make(map[string]Node)
	}

	sanitizedPath, vars := sanitizePath(path)
	node := Node{
		handler:   handler,
		variables: vars,
	}
	// add a node instead, to store metadata about the paramters
	handlers[method][sanitizedPath] = node
}

package golem

import (
	"github.com/julienschmidt/httprouter"
)

func New() *Router {
	return &Router{
		InnerRouter: httprouter.New(),
	}
}

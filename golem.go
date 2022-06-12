package golem

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wxlai90/golem/router"
)

func New() *router.Router {
	return &router.Router{
		InnerRouter: httprouter.New(),
	}
}

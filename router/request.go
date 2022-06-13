package router

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Request struct {
	R       *http.Request
	Cookies map[string]string
	Params  map[string]string
	Query   map[string]string
	Bag     Bag
}

type Bag struct {
	bag map[string]interface{}
}

func NewRequest(r *http.Request, p httprouter.Params) *Request {
	req := Request{
		R: r,
	}

	req.initBag()
	req.parseCookies()
	req.parseQueries()
	req.parseParams(p)

	return &req
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

func (r *Request) parseQueries() {
	r.Query = map[string]string{}
	values := r.R.URL.Query()
	for k, v := range values {
		r.Query[k] = v[0]
	}
}

func (r *Request) initBag() {
	r.Bag = Bag{
		bag: map[string]interface{}{},
	}
}

func (r *Request) Put(key string, value interface{}) {
	// ignore possibility of overwriting
	r.Bag.bag[key] = value
}

func (r *Request) Get(key string) (interface{}, error) {
	if value, ok := r.Bag.bag[key]; ok {
		return value, nil
	}

	return nil, errors.New("Key not found.")
}

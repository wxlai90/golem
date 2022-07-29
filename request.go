package golem

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Body struct {
	RawBytes []byte
}

func (b *Body) Unmarshal(m interface{}) error {
	return json.Unmarshal(b.RawBytes, m)
}

type Request struct {
	*http.Request
	Cookies    map[string]string
	Params     map[string]string
	Query      map[string]string
	Bag        Bag
	Body       *Body
	FormValues map[string]string
}

type Bag struct {
	bag map[string]interface{}
}

func NewRequest(r *http.Request, p httprouter.Params) *Request {
	req := Request{
		Request: r,
	}

	req.initBag()
	req.parseCookies()
	req.parseQueries()
	req.parseParams(p)
	req.ParseForm()
	req.parseRequestBody()

	return &req
}

func (r *Request) parseCookies() {
	r.Cookies = map[string]string{}
	for _, cookie := range r.Request.Cookies() {
		r.Cookies[cookie.Name] = cookie.Value
	}
}

func (r *Request) parseParams(params httprouter.Params) {
	r.Params = map[string]string{}
	for _, param := range params {
		r.Params[param.Key] = param.Value
	}
}

func (r *Request) parseForm() {
	MAX_MEMORY := int64(32 << 20)
	r.ParseMultipartForm(MAX_MEMORY)

	for k, v := range r.Form {
		r.FormValues[k] = v[0]
	}
}

func (r *Request) parseQueries() {
	r.Query = map[string]string{}
	values := r.Request.URL.Query()
	for k, v := range values {
		r.Query[k] = v[0]
	}
}

func (r *Request) parseRequestBody() {
	body, err := io.ReadAll(r.Request.Body)
	if err != nil {
		log.Panicln(err)
	}
	defer r.Request.Body.Close()

	r.Body = &Body{
		RawBytes: body,
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

func (r *Request) Get(key string) (interface{}, bool) {
	value, ok := r.Bag.bag[key]
	return value, ok
}

package golem

import (
	"encoding/json"
	"io"
	"net/http"
)

type Body struct {
	request *http.Request
}

func (b *Body) Unmarshal(m interface{}) error {
	body := b.request.Body
	bs, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	defer b.request.Body.Close()

	return json.Unmarshal(bs, m)
}

type Request struct {
	*http.Request
	Query map[string]string
	Body  *Body
}

func NewRequest(r *http.Request) *Request {
	req := Request{
		Request: r,
		Body: &Body{
			request: r,
		},
	}
	req.parseQueries()

	return &req
}

func (r *Request) GetCookie(name string) string {
	for _, cookie := range r.Request.Cookies() {
		if cookie.Name == name {
			return cookie.Value
		}
	}

	return ""
}

func (r *Request) GetQuery(name string) string {
	if val, ok := r.Query[name]; ok {
		return val
	}

	return ""
}

func (r *Request) parseQueries() {
	r.Query = map[string]string{}
	values := r.URL.Query()
	for k, v := range values {
		r.Query[k] = v[0]
	}
}

func (r *Request) Params(name string) string {
	return r.PathValue(name)
}

func Decode[T any](data []byte) (T, error) {
	var t T
	err := json.Unmarshal(data, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

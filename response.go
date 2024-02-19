package golem

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	res := Response{
		ResponseWriter: w,
	}

	return &res
}

func (r Response) Json(response interface{}) {
	r.Header().Set("Content-Type", "application/json")
	json.NewEncoder(r).Encode(response)
}

func (r Response) Send(response string) {
	io.WriteString(r, response)
}

func (r Response) Cookie(key, value string) {
	http.SetCookie(r, &http.Cookie{
		Name:  key,
		Value: value,
	})
}

func (r Response) Status(statusCode int) Response {
	r.WriteHeader(statusCode)
	return r
}

package router

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	W http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	res := Response{
		W: w,
	}

	return &res
}

func (r Response) JSON(response interface{}) {
	r.W.Header().Set("Content-Type", "application/json")
	json.NewEncoder(r.W).Encode(response)
}

func (r Response) Send(response string) {
	io.WriteString(r.W, response)
}

func (r Response) Cookie(key, value string) {
	http.SetCookie(r.W, &http.Cookie{
		Name:  key,
		Value: value,
	})
}

func (r Response) Status(statusCode int) Response {
	r.W.WriteHeader(statusCode)
	return r
}

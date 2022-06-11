package router

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	W http.ResponseWriter
}

func (r Response) JSON(response interface{}) {
	json.NewEncoder(r.W).Encode(response)
}

func (r Response) Send(response string) {
	io.WriteString(r.W, response)
}

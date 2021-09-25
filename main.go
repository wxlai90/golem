package main

import (
	"golem/router"
	"io"
	"net/http"
)

func main(){
	r := router.New()
	
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "GET: Hello World")
	})

	r.POST("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "POST: Hello Again World!")
	})

	http.ListenAndServe(":5000", r)
}
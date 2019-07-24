package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "woof woof")
}

type hotcat int

func (m hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "meow")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/dog/", d) // path: /dog/catches/any/path
	http.Handle("/cat", c)  // catches only exact match path: /cat

	http.ListenAndServe(":8080", nil)
}

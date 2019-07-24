package main

import (
	"io"
	"net/http"
)

type hotdog int

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "woof woof")
}

type hotcat int

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "meow")
}

func main() {

	http.Handle("/dog/", http.HandlerFunc(d)) // path: /dog/catches/any/path
	http.Handle("/cat", http.HandlerFunc(c))  // catches only exact match path: /cat

	http.ListenAndServe(":8080", nil)
}

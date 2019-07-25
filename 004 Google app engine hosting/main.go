package main

import (
	"net/http"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir(".")))
}

// no func main()

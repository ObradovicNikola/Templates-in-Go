package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
}

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	tmpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

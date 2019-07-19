package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	result := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}
	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", result)
	if err != nil {
		log.Fatalln(err)
	}
}

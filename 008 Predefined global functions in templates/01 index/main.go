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

	array := []string{"zero", "one", "two", "three", "four", "five"}

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", array)
	if err != nil {
		log.Fatalln(err)
	}
}

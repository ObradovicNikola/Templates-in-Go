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
	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}

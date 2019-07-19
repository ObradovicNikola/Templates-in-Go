package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	err := tmpl.ExecuteTemplate(os.Stdout, "default.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}

// go run main.go > index.html

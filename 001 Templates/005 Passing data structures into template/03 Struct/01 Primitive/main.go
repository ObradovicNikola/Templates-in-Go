package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type car struct {
	Brand  string
	Model  string
	Engine int // hp
}

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	spyder := car{
		Brand:  "Porsche",
		Model:  "Spyder 918",
		Engine: 600,
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", spyder)
	if err != nil {
		log.Fatalln(err)
	}
}

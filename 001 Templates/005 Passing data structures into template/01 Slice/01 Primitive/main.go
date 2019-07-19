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

	cars := []string{"Porsche", "Mazda", "Peugeot", "BMW"}

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", cars)
	if err != nil {
		log.Fatalln(err)
	}
}

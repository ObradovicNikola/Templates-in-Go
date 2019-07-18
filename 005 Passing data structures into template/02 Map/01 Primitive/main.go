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

	capitals := map[string]string{
		"Austria": "Vienna",
		"Belgium": "Brussels",
		"Serbia":  "Belgrade",
		"Finland": "Helsinki",
		"France":  "Paris",
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", capitals)
	if err != nil {
		log.Fatalln(err)
	}
}

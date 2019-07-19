package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseFiles("tmpl.gohtml")
	if err != nil {
		log.Fatal("error creating file", err)
	}

	err = tmpl.Execute(os.Stdout, 42)
	if err != nil {
		log.Fatalln(err)
	}
}

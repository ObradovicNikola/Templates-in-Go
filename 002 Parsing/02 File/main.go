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

	f, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer f.Close()

	err = tmpl.Execute(f, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tmpl *template.Template

func double(x float64) float64 {
	return x + x
}

func square(x float64) float64 {
	return math.Pow(x, 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"double": double,
	"square": square,
	"sqRoot": sqRoot,
}

func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.gohtml"))
}

func main() {

	magicNumber := 9

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", float64(magicNumber))
	if err != nil {
		log.Fatalln(err)
	}
}

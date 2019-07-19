package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tmpl *template.Template

type car struct {
	Brand  string
	Model  string
	Engine int // hp
}

/*
	uc - is what the func will be called in the template
	uc - is the ToUpper func from package strings
	ft - isa func I declared
	ft - slices a string, returning the first three characters
*/
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {

	spyder := car{
		Brand:  "Porsche",
		Model:  "Spyder 918",
		Engine: 600,
	}

	aventador := car{
		Brand:  "Lamborghini",
		Model:  "Aventador S",
		Engine: 730,
	}

	veyron := car{
		Brand:  "Bugatti",
		Model:  "Veyron Grand Sport",
		Engine: 987,
	}

	corvette := car{
		Brand:  "Chevrolet",
		Model:  "Corvette ZR1",
		Engine: 755,
	}

	cars := []car{spyder, aventador, veyron, corvette}

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", cars)
	if err != nil {
		log.Fatalln(err)
	}
}

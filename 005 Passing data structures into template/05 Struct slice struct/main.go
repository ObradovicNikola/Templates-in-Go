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

	race := struct {
		Track string
		Date  string
		Cars  []car
	}{
		Track: "Miseluk",
		Date:  "18/7/2019",
		Cars:  cars,
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", race)
	if err != nil {
		log.Fatalln(err)
	}
}

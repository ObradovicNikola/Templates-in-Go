package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var tmpl *template.Template

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

	/*
		"01/02 03:04:05PM '06 -0700"
		"MM/DD TIME  TIME YEAR  MST"   // MST is GMT -0700
		each number represent specific variable in date
	*/

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", time.Now().Format("02.01.2006. 03:04PM"))
	if err != nil {
		log.Fatalln(err)
	}
}

//functionMap := FuncMap()

//Funcs(functionMap)

package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"str1": strings.ToUpper,
	"str2": strings.ToLower,
}

type Person struct {
	Name    string
	Grade   string
	Pronoun string
}

func main() {

	roll1 := Person{
		Name:    "Abdul Console",
		Grade:   "A1",
		Pronoun: "his",
	}
	roll2 := Person{
		Name:    "John Mary",
		Grade:   "B2",
		Pronoun: "her",
	}
	roll3 := Person{
		Name:    "James Tony",
		Grade:   "C4",
		Pronoun: "his",
	}
	roll4 := Person{
		Name:    "Ibrahim Okey",
		Grade:   "A+",
		Pronoun: "his",
	}
	roll5 := Person{
		Name:    "Ali Aduke",
		Grade:   "B+",
		Pronoun: "her",
	}

	persons := []Person{roll1, roll2, roll3, roll4, roll5}
	//template1, err := template.ParseFiles("template.gohtml")
	template1, err := template.New("template.gohtml").Funcs(fm).ParseFiles("template.gohtml")

	if err != nil {
		log.Fatal(err)
		return
	}
	template1.ExecuteTemplate(os.Stdout, "template.gohtml", persons)
}

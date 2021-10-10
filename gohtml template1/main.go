package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	var slice1 = []string{"John", "Mark", "Elsa", "Sharlock", "Emma", "Harmo"}
	templ, err := template.ParseFiles("template.gohtml")

	if err != nil {
		log.Fatal(err)
		return
	}
	templ.ExecuteTemplate(os.Stdout, "template.gohtml", slice1)
}

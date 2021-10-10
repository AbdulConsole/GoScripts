package main

import (
	"log"
	"os"
	"text/template"
)

//Html-tise content of template

func main() {

	var slice1 = []string{"Ade", "Ola", "Ayo", "Odun", "Ire", "Osu"}
	templ, err := template.ParseFiles("template.gohtml")

	if err != nil {
		log.Fatal(err)
		return
	}
	templ.ExecuteTemplate(os.Stdout, "template.gohtml", slice1)
}

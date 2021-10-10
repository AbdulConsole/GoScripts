package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	var var1 = "Hi world!"
	templ, err := template.ParseFiles("template.gohtml")

	if err != nil {
		log.Fatal(err)
		return
	}
	templ.ExecuteTemplate(os.Stdout, "template.gohtml", var1)
}

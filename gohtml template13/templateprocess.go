package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tmpl, err := template.ParseGlob("templates/*.gohtml")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", nil)

}

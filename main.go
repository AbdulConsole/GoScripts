package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	templ, err := template.ParseFiles("template.gohtml")

	if err != nil {
		log.Fatal(err)
		return
	}

	templ.Execute(os.Stdout, nil)
}

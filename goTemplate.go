package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name    string
	Hobbies []string
}

const data = "The name of the person is {{.Name}}.{{range .Hobbies}}He likes {{.}}{{end}}"

func main() {

	person := Person{
		Name:    "Abdul",
		Hobbies: []string{"Gisting", "Hacking"},
	}

	templ := template.New("test")

	templ, err := templ.Parse(data)

	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	templ.Execute(os.Stdout, person)
}

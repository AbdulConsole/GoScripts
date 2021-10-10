package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name string
}

func main() {
	name := Person{"Abdul Console"}

	template1 := template.New("test")

	template1, err := template1.Parse("Hello World {{.Name}}")

	if err != nil {
		log.Fatal(err)
	}
	template1.Execute(os.Stdout, name)
}

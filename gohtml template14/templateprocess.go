package main

import (
	"log"
	"os"
	"text/template"
)

// go run templateprocess.go > web.html
// type Person string

// func (p Person) func1() string {
// 	return "There we have some text" + string(p)
// }
func main() {
	var var1 = "Coming From Process File"
	tmpl, err := template.ParseGlob("templates/*.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", var1)
}

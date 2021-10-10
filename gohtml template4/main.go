package main

import (
	"log"
	"os"
	"text/template"
)

//Html-tise map content of template

func main() {

	// var slice1 = []string{"Ade", "Ola", "Ayo", "Odun", "Ire", "Osu"}
	var map1 = map[string]string{
		"Abia":    "Umaiyah",
		"Oyo":     "Ibadan",
		"Kwara":   "Ilorin",
		"Osun":    "Oshogo",
		"Ogun":    "Abeokuta",
		"Lagos":   "Ikeja",
		"Nigeria": "Abuja",
	}

	templ, err := template.ParseFiles("template.gohtml")

	if err != nil {
		log.Fatal(err)
		return
	}
	templ.ExecuteTemplate(os.Stdout, "template.gohtml", map1)
}

//functionMap := FuncMap()

//Funcs(functionMap)

package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Name            string
	Grade           string
	Pronoun         string
	Member          bool
	Temporarymember bool
	Num1            int
	Num2            int
}

func customStringFunc(strPara string) string {
	strPara = strings.TrimLeft(strPara, "Ali")
	//strPara = strPara[:2]
	return strPara
}
func main() {

	roll1 := Person{
		Name:            "Abdul Console",
		Grade:           "A1",
		Pronoun:         "his",
		Member:          false,
		Temporarymember: true,
		Num1:            8,
		Num2:            4,
	}
	roll2 := Person{
		Name:            "John Mary",
		Grade:           "B2",
		Pronoun:         "her",
		Member:          true,
		Temporarymember: false,
	}
	roll3 := Person{
		Name:            "James Tony",
		Grade:           "C4",
		Pronoun:         "his",
		Member:          false,
		Temporarymember: true,
	}
	roll4 := Person{
		Name:            "Ibrahim Okey",
		Grade:           "A+",
		Pronoun:         "his",
		Member:          false,
		Temporarymember: false,
	}
	roll5 := Person{
		Name:            "Ali Aduke",
		Grade:           "B+",
		Pronoun:         "her",
		Member:          true,
		Temporarymember: false,
	}

	persons := []Person{roll1, roll2, roll3, roll4, roll5}

	var fm = template.FuncMap{
		"str1": strings.ToUpper,
		"str2": strings.ToLower,
		"str3": customStringFunc,
	}

	//template1, err := template.ParseFiles("template.gohtml")
	template1, err := template.New("template.gohtml").Funcs(fm).ParseFiles("template.gohtml")

	if err != nil {
		log.Fatal(err)
		return
	}
	template1.ExecuteTemplate(os.Stdout, "template.gohtml", persons)
}

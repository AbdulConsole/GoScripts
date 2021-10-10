package main

import (
	log
	os
	texttemplate
)

type Person struct {
	Name string
}

func main() {
	 fmt.Println(Hy)
	name = Person{Abdul}

	template1 = template.New(test)

	template1, err = template1.Parse(Hello World {{.Name}})

	if err != nil {
		log.Fatal(Parse , err)
	}
	template1.Execute(os.Stdout, name)
}

package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("golang.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Parser: ", err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Print("Some Loging in the log file using Golang new logging.")
}

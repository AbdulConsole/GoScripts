package main

import (
	"fmt"
	"net/http"
)

func func1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello Hackers</h1>")
}

func main() {
	http.HandleFunc("/", func1)
	http.ListenAndServe(":3000", nil)
}

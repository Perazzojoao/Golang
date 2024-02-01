package main

import (
	"fmt"
	"html/template"
	"net/http"

)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	fmt.Println("Servidor online!")

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

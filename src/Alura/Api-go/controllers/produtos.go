package controllers

import (
	"html/template"
	"net/http"

	"api-web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	listaProdutos := models.SelectAllProducts()

	temp.ExecuteTemplate(w, "Index", listaProdutos)
}

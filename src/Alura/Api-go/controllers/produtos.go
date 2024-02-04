package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"api-web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	listaProdutos := models.SelectAllProducts()

	temp.ExecuteTemplate(w, "Index", listaProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		err := errors.New("ERROR: You suposed to meke a POST request.")
		fmt.Println(err.Error())
		return
	}

	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	estoque := r.FormValue("estoque")

	precoConvertido, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		log.Println("Error trying to convert 'preco':", err.Error())
		return
	}

	estoqueConvertido, err := strconv.Atoi(estoque)
	if err != nil {
		log.Println("Error trying to convert 'estoque'", err.Error())
		return
	}
	models.AddNewProduct(nome, descricao, precoConvertido, estoqueConvertido)
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	produto := models.GetProduto(r.URL.Query().Get("id"))

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		err := errors.New("ERROR: You suposed to meke a PUT request.")
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", 401)
		return
	}

	id := r.FormValue("id")
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	estoque := r.FormValue("estoque")

	precoConvertido, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		log.Println("Error trying to convert 'preco':", err.Error())
		return
	}
	estoqueConvertido, err := strconv.Atoi(estoque)
	if err != nil {
		log.Println("Error trying to convert 'estoque'", err.Error())
		return
	}

	models.UpdateProduto(id, nome, descricao, precoConvertido, estoqueConvertido)
	http.Redirect(w, r, "/", 301)
}

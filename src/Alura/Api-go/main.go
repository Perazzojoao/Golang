package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"api-web/produtos"
)

func conectarDb() *sql.DB {
	db, err := sql.Open("mysql", "root:joao1998J@tcp(localhost:3306)/alura_loja")

	if err != nil {
		fmt.Println("ERROR -> Fail to validate sql.Open() arguments")
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("ERROR -> Fail to verify connection with db.Ping()")
		panic(err.Error())
	}
	return db
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectarDb()
	defer db.Close()

	fmt.Println("Servidor online!")

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	listaProdutos := produtos.NewProduto("Camiseta", "Confortável", 89, 3)
	listaProdutos = produtos.NewProduto("Tênis", "Confortável", 39, 12)
	listaProdutos = produtos.NewProduto("Fone", "Muito bom", 59, 6)
	listaProdutos = produtos.NewProduto("Notebook", "Muito rápido", 3688.99, 22)

	temp.ExecuteTemplate(w, "Index", listaProdutos)
}

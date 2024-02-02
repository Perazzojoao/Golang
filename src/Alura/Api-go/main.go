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
	fmt.Println("Servidor online!")

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectarDb()
	defer db.Close()

	selectAll, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		fmt.Println("ERROR -> Unable to perform a query")
		fmt.Println(err.Error())
	}

	var p produtos.Produto
	listaProdutos := make([]produtos.Produto, 0, 5)

	for selectAll.Next() {
		var id, estoque int
		var nome, descricao string
		var preco float64

		err := selectAll.Scan(&id, &nome, &descricao, &preco, &estoque)
		if err != nil {
			fmt.Println("ERROR -> Unable to scan each table line")
			fmt.Println(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Estoque = estoque

		listaProdutos = append(listaProdutos, p)
	}

	temp.ExecuteTemplate(w, "Index", listaProdutos)
}

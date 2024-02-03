package models

import (
	"fmt"

	"api-web/db"

)

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
	Estoque   int
}

func SelectAllProducts() []Produto {
	db := db.ConectarDb()
	defer db.Close()

	selectAll, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		fmt.Println("ERROR -> Unable to perform a query")
		fmt.Println(err.Error())
		return make([]Produto, 0, 3)
	}

	var p Produto
	listaProdutos := make([]Produto, 0, 5)

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
	return listaProdutos
}

func AddNewProduct(nome string, descricao string, preco float64, estoque int) {
	db := db.ConectarDb()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO alura_loja.produtos (nome, descricao, preco, estoque) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println("ERROR -> Wrong insert arguments.")
		fmt.Println(err.Error())
		return
	}
	stmt.Exec(nome, descricao, preco, estoque)
}

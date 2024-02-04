package models

import (
	"fmt"
	"log"

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

	selectAll, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
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
	stmt.Close()
	log.Printf("Product '%s' added to data base.\n", nome)
}

func DeletaProduto(id string) {
	db := db.ConectarDb()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM alura_loja.produtos WHERE id = ?")
	if err != nil {
		fmt.Println("ERROR -> Wrong delete arguments.")
		fmt.Println(err.Error())
		return
	}
	stmt.Exec(id)
	stmt.Close()

	log.Println("Product ID =", id, "deleted.")
}

func GetProduto(idProduto string) Produto {
	db := db.ConectarDb()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM alura_loja.produtos WHERE id = ?")
	if err != nil {
		fmt.Println("ERROR -> wrong select arguments.")
		fmt.Println(err.Error())
		return Produto{}
	}
	product, err := stmt.Query(idProduto)
	if err != nil {
		fmt.Println("ERROR -> Unable to perform the query.")
		fmt.Println(err.Error())
		return Produto{}
	}
	stmt.Close()

	var id, estoque int
	var nome, descricao string
	var preco float64

	product.Next()
	err = product.Scan(&id, &nome, &descricao, &preco, &estoque)
	if err != nil {
		fmt.Println("ERROR -> Unable to scan table line")
		fmt.Println(err.Error())
	}

	return Produto{id, nome, descricao, preco, estoque}
}

func UpdateProduto(id string, nome string, descricao string, preco float64, estoque int) {
	db := db.ConectarDb()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE alura_loja.produtos SET nome = ?, descricao = ?, preco = ?, estoque = ? WHERE id = ?")
	if err != nil {
		fmt.Println("ERROR -> wrong update arguments.")
		fmt.Println(err.Error())
		return
	}

	stmt.Exec(nome, descricao, preco, estoque, id)
	stmt.Close()
	log.Println("Produto id =", id, "updated.")
}

package produtos

import "slices"

type Produto struct {
	Nome      string
	Descricao string
	Preco     float64
	Estoque   int
}

var produtos = []Produto{}

func NewProduto(nome string, descricao string, preco float64, estoque int) *[]Produto {
	if slices.Contains[[]Produto](produtos, Produto{nome, descricao, preco, estoque}) {
		return &produtos
	}

	produtos = append(produtos, Produto{nome, descricao, preco, estoque})

	return &produtos
}

func (p *Produto) GetNome() string {
	return p.Nome
}

func (p *Produto) GetDescricao() string {
	return p.Descricao
}

func (p *Produto) GetPreco() float64 {
	return p.Preco
}

func (p *Produto) GetEstoque() int {
	return p.Estoque
}

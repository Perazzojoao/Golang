package produtos

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
	Estoque   int
}

func NewProduto(nome string, descricao string, preco float64, estoque int) *Produto {
	return &Produto{0, nome, descricao, preco, estoque}
}

func NewListaProduto(listaProdutos ...Produto) *[]Produto {
	produtos := []Produto{}
	for _, produto := range listaProdutos {
		produtos = append(produtos, produto)
	}
	return &produtos
}

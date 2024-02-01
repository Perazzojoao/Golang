package clientes

type Titular struct {
	nome     string
	cpf      int
	profisao string
}

func NewTitular(nome string, cpf int, profissao string) *Titular {
	return &Titular{nome, cpf, profissao}
}

func (t *Titular) Profisao() string {
	return t.profisao
}

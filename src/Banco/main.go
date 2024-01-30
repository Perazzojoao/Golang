package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) sacar(valorSaque float64) {
	podeSacar := valorSaque <= c.saldo && valorSaque > 0

	if podeSacar {
		c.saldo -= valorSaque
		fmt.Println("Saque realizado com sucesso!")
	} else {
		fmt.Println("Saldo insuficiente")
	}
	println()
}

func main() {
	cliente01 := ContaCorrente{
		titular:       "João Victor",
		numeroAgencia: 552,
		numeroConta:   3301,
		saldo:         152.34,
	} // Não é obrigatório atribuir valores a todos os campos

	cliente02 := ContaCorrente{"Guilherme", 442, 3302, 3242.52} // Obrigatório preencher todos os campos

	fmt.Println(cliente01)
	fmt.Println(cliente02)
}

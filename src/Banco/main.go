package main

import (
	"fmt"

	"banco/contas"
)

func main() {
	cliente01 := contas.ContaCorrente{
		Titular:       "João Victor",
		NumeroAgencia: 552,
		NumeroConta:   3301,
		Saldo:         152.34,
	} // Não é obrigatório atribuir valores a todos os campos

	cliente02 := contas.ContaCorrente{"Guilherme", 442, 3302, 3242.52} // Obrigatório preencher todos os campos

	status, err := cliente02.Transferir(-1000, &cliente01)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
	fmt.Println(cliente01)
	fmt.Println(cliente02)
}

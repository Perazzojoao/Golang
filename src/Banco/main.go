package main

import (
	"fmt"

	"banco/clientes"
	"banco/contas"

)

func main() {
	clienteBruno := clientes.Titular{
		Nome:      "Bruno",
		Cpf:       70142572477,
		Proffisao: "Programador",
	}
	contaBruno := contas.ContaCorrente{
		Titular:       clienteBruno,
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}

	fmt.Println(contaBruno)
}

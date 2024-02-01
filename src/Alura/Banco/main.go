package main

import (
	"fmt"

	"banco/clientes"
	"banco/contas"
)

func main() {
	cliente1 := clientes.NewTitular("João Victor", 70242452477, "Desenvolvedor")
	conta1 := contas.NewContaCorrente(*cliente1, 123, 123456, 11000)

	fmt.Println("Saldo inicial:", conta1.GetSaldo())

	contas.PagarBoleto(conta1, 1000)

	fmt.Println("Saldo final:", conta1.GetSaldo())
}

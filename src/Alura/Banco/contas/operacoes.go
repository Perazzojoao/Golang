package contas

import "fmt"

type conta interface {
	GetSaldo() float64
	Sacar(valorSaque float64)
	Depositar(valor float64) error
}

func PagarBoleto(contaUtilizada conta, valor float64) {
	contaUtilizada.Sacar(valor)

	fmt.Println("Boleto pago com sucesso!")
}

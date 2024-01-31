package contas

import (
	"errors"
	"fmt"

)

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) {
	podeSacar := valorSaque <= c.Saldo && valorSaque > 0

	if podeSacar {
		c.Saldo -= valorSaque
		fmt.Println("Saque realizado com sucesso!")
	} else {
		fmt.Println("Saldo insuficiente")
	}
	println()
}

func (c *ContaCorrente) Depositar(valor float64) error {
	if valor > 0 {
		c.Saldo += valor
		return nil
	}
	return fmt.Errorf("Não foi possível realizar o depósito.")
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) (bool, error) {
	if valor < c.Saldo && valor > 0 {
		c.Saldo -= valor
		contaDestino.Saldo += valor
		return true, nil
	}
	return false, errors.New("Não foi possível realizar a transferência.")
}

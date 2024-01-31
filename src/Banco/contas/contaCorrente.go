package contas

import (
	"errors"
	"fmt"

	c "banco/clientes"
)

type ContaCorrente struct {
	Titular       c.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) {
	podeSacar := valorSaque <= c.saldo && valorSaque > 0

	if podeSacar {
		c.saldo -= valorSaque
		fmt.Println("Saque realizado com sucesso!")
	} else {
		fmt.Println("Saldo insuficiente")
	}
	println()
}

func (c *ContaCorrente) Depositar(valor float64) error {
	if valor > 0 {
		c.saldo += valor
		return nil
	}
	return fmt.Errorf("Não foi possível realizar o depósito.")
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) (bool, error) {
	if valor < c.saldo && valor > 0 {
		c.saldo -= valor
		contaDestino.saldo += valor
		return true, nil
	}
	return false, errors.New("Não foi possível realizar a transferência.")
}

func (c *ContaCorrente) getSaldo() float64 {
	return c.saldo
}

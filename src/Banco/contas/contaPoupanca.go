package contas

import (
	"fmt"

	c "banco/clientes"

)

type ContaPoupanca struct {
	titular       c.Titular
	numeroAgencia int
	numeroConta   int
	operacao      int
	saldo         float64
}

func NewContaPoupanca(titular c.Titular, numAgencia int, numConta int, operacao int, saldoInicial float64) *ContaPoupanca {
	o := ContaPoupanca{titular, numAgencia, numConta, operacao, saldoInicial}
	return &o
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) Sacar(valorSaque float64) {
	podeSacar := valorSaque <= c.saldo && valorSaque > 0

	if podeSacar {
		c.saldo -= valorSaque
		fmt.Println("Saque realizado com sucesso!")
	} else {
		fmt.Println("Saldo insuficiente")
	}
	println()
}

func (c *ContaPoupanca) Depositar(valor float64) error {
	if valor > 0 {
		c.saldo += valor
		return nil
	}
	return fmt.Errorf("Não foi possível realizar o depósito.")
}

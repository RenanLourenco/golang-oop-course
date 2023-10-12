package contas

import (
	"fmt"

	"banco.com/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!", c.saldo
	} else {
		return "saldo insuficiente", c.saldo
	}

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
	} else {
		fmt.Println("Valor incorreto")
	}
}

func (c *ContaCorrente) VisualizarSaldo() float64 {
	return c.saldo
}

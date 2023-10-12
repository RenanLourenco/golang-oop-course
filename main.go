package main

import (
	"fmt"

	"banco.com/clientes"
	"banco.com/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	clienteRenan := clientes.Titular{Nome: "Renan", CPF: "35796035827", Profissao: "Desenvolvedor"}

	contaDoRenan := contas.ContaCorrente{Titular: clienteRenan, NumeroAgencia: 123, NumeroConta: 321}
	contaDoRenan.Depositar(3000)

	contaPoupancaRenan := contas.ContaPoupanca{Titular: clienteRenan, NumeroAgencia: 123, NumeroConta: 321, Operacao: 3}
	contaPoupancaRenan.Depositar(300)


	PagarBoleto(&contaPoupancaRenan, 200)

	saldo := contaPoupancaRenan.VisualizarSaldo()

	fmt.Println(saldo)

}

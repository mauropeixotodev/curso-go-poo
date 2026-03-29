package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagamentoDeBoleto(conta verificarConta, valorDoBoleto float64) string {
	if valorDoBoleto > 0 && valorDoBoleto <= conta.ObterSaldo() {
		conta.Sacar(valorDoBoleto)
		return "Pagamento realizado com sucesso"
	} else {
		return "Saldo insuficiente para realizar o pagamento"
	}

}

type verificarConta interface {
	ObterSaldo() float64
	Sacar(valorDoSaque float64) string
}

func main() {

	contaDoDenis := contas.ContaPoupanca{
		Titular: clientes.Titular{
			Nome:      "Denis",
			Cpf:       "987.654.321-00",
			Profissao: "Designer"},
	}

	ContaDaPati := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Pati",
			Cpf:       "123.456.789-00",
			Profissao: "Programadora"},
	}
	fmt.Println("Conta do Denis:", contaDoDenis)
	fmt.Println("Conta da Pati:", ContaDaPati)

	contaDoDenis.Depositar(1000)
	fmt.Println("Saldo do Denis após depósito:", contaDoDenis.ObterSaldo())
	ContaDaPati.Depositar(500)
	fmt.Println("Saldo da Pati após depósito:", ContaDaPati.ObterSaldo())

	PagamentoDeBoleto(&contaDoDenis, 200)
	fmt.Println("Saldo do Denis após pagamento de boleto:", contaDoDenis.ObterSaldo())

}

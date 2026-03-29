package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	contaDoBruno := contas.ContaCorrente{
		Titular:       clientes.Titular{Nome: "Bruno", Cpf: "123.456.789-00", Profissao: "Desenvolvedor"},
		NumeroAgencia: 123,
		NumeroConta:   456,
	}

	fmt.Println("Conta do Bruno:", contaDoBruno)
	fmt.Println("Saldo do Bruno:", contaDoBruno.ObterSaldo())

	contaDoBruno.Depositar(500.0)
	fmt.Println("Saldo do Bruno após depósito:", contaDoBruno.ObterSaldo())

}

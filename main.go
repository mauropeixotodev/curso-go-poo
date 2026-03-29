package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente para realizar o saque"
	}
}

func (c ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito <= 0 {
		return "Valor do depósito deve ser maior que zero", c.saldo
	}
	c.saldo += valorDoDeposito
	return "Depósito realizado com sucesso", c.saldo
}

func main() {

	contaDoSilvia := ContaCorrente{}
	contaDoSilvia.titular = "Silvia"
	contaDoSilvia.numeroAgencia = 863
	contaDoSilvia.numeroConta = 863452
	contaDoSilvia.saldo = 1000.00

	fmt.Println(contaDoSilvia.Sacar(300))
	status, saldo := contaDoSilvia.Depositar(500)
	fmt.Println(status)
	fmt.Printf("Saldo atual: R$ %.2f\n", saldo)

	fmt.Println("contado do Silvia tem saldo de R$", contaDoSilvia.saldo)

}

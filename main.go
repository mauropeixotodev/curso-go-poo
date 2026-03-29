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

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito <= 0 {
		return "Valor do depósito deve ser maior que zero", c.saldo
	}
	c.saldo += valorDoDeposito
	return "Depósito realizado com sucesso", c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}

}

func main() {

	contaDoSilvia := ContaCorrente{titular: "Silvia", saldo: 300}
	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 100}
	status := contaDoSilvia.Transferir(200, &contaDoGustavo)
	fmt.Printf("Status da transferencia: %t\n", status)
	fmt.Printf("Saldo da conta do Silvia: %.2f\n", contaDoSilvia.saldo)
	fmt.Printf("Saldo da conta do Gustavo: %.2f\n", contaDoGustavo.saldo)

}

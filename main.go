package main

import (
	c "banco/contas"
	"fmt"
)

func main() {

	contaDoSilvia := c.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := c.ContaCorrente{Titular: "Gustavo", Saldo: 100}
	status := contaDoSilvia.Transferir(200, &contaDoGustavo)
	fmt.Printf("Status da transferencia: %t\n", status)
	fmt.Printf("Saldo da conta do Silvia: %.2f\n", contaDoSilvia.Saldo)
	fmt.Printf("Saldo da conta do Gustavo: %.2f\n", contaDoGustavo.Saldo)

}

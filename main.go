package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaDoMauro := ContaCorrente{
		titular:       "Mauro Silva",
		numeroAgencia: 4321,
		numeroConta:   98765,
		saldo:         2500.75,
	}

	var contaDoCristina *ContaCorrente
	contaDoCristina = new(ContaCorrente)
	contaDoCristina.titular = "Cristina Oliveira"
	contaDoCristina.numeroAgencia = 4321
	contaDoCristina.numeroConta = 98766
	contaDoCristina.saldo = 3000.00
	fmt.Println(contaDoMauro)
	fmt.Println(*contaDoCristina)

	var titular string = "João da Silva"
	var numeroAgencia int = 1234
	var numeroConta int = 56789
	var saldo float64 = 1000.50

	fmt.Printf("Titular: %s\n", titular)
	fmt.Printf("Agência: %d\n", numeroAgencia)
	fmt.Printf("Conta: %d\n", numeroConta)
	fmt.Printf("Saldo: %.2f\n", saldo)
}

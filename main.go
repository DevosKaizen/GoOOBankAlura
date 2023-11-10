package main

import (
	"fmt"

	"banco/clientes"
	"banco/contas"
)

func main() {

	fmt.Println()
	// USANDO NEW NAO FUNCIONOU
	// contaDaSilvia :=  new(ContaCorrente) // ou ContaCorrente{}
	// contaDaSilvia.titular = "Silvia"
	// contaDaSilvia.saldo = 300
	// fmt.Println(contaDaSilvia)

	// contaDoGustavo := new(ContaCorrente)
	// contaDoGustavo.titular = "Gustavo"
	// contaDoGustavo.saldo = 100

	contaDaSilvia := contas.ContaCorrente{Titular: clientes.Titular{}, Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: clientes.Titular{}, Saldo: 100}

	fmt.Println(contaDoGustavo)
	// usar & comercial pro ponteiro
	status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	fmt.Println(status)
	fmt.Println("o novo Saldo do Gustavo é", contaDoGustavo.Saldo)
	fmt.Println("o novo Saldo da Silvia é", contaDaSilvia.Saldo)

}

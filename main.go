package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	fmt.Println()

	contaDaSilvia :=  new(ContaCorrente) // ou ContaCorrente{} 
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500
	fmt.Println(contaDaSilvia)

	// valorDoSaque := 200.
	


	fmt.Println(contaDaSilvia.Sacar(200))
	fmt.Println("O valor do SALDO é: ", contaDaSilvia.saldo)


} //equivalente a init, self/ 'c' (poderia ser qualquer nome mas em go é a primeira letra do ponteiro) é o endereço de *ContaCorrente
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso! "
	} else {
		return "Saldo insuficiente! "
	}
}

package main

import "fmt"

func main() {

	fmt.Println()
	banco()
	

}	//equivalente a init, self/ 'c' (poderia ser qualquer nome mas em go é a primeira letra do ponteiro) é o endereço de *ContaCorrente
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo
	if podeSacar{
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso! "
	} else {
		return "Saldo insuficiente! "
	}
}
func banco(Sacar()) {
	type ContaCorrente struct {
		titular       string
		numeroAgencia int
		numeroConta   int
		saldo         float64
	}
	contaDoGuilherme := ContaCorrente{titular:"Guilherme", numeroAgencia: 5473, numeroConta: 326545, saldo: 200}
	fmt.Println(contaDoGuilherme)
	
	contaDaCris := new(ContaCorrente)
	fmt.Println(contaDaCris)

	contaDaSilvia := new(ContaCorrente)
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500
	fmt.Println(contaDaSilvia)

	valorDoSaque := 200.
	fmt.Println("O valor do saque é: ",valorDoSaque)
	fmt.Println(contaDaSilvia.Sacar(300))
}


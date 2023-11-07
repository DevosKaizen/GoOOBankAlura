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

	status, valor := contaDaSilvia.Depositar(20.)
	fmt.Println("O valor do SALDO é: ", contaDaSilvia.saldo)

	fmt.Println(status, valor)

} //equivalente a init, self/ 'c' (poderia ser qualquer nome mas em go é a primeira letra do ponteiro) é o endereço de *ContaCorrente // ENtrada   //retorno
func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 //&& valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo += valorDoSaque
		return "Deposito realizado com sucesso! ", c.saldo
	} else {
		return "Saldo insuficiente! ", c.saldo
	}
}
func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64){
	
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Saque realizado com sucesso! seu novo saldo é ", c.saldo
	} else {
		return "Saldo insuficiente! seu saldo é ", c.saldo
	}
}
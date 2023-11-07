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
	// USANDO NEW NAO FUNCIONOU
	// contaDaSilvia :=  new(ContaCorrente) // ou ContaCorrente{} 
	// contaDaSilvia.titular = "Silvia"
	// contaDaSilvia.saldo = 300
	// fmt.Println(contaDaSilvia)

	// contaDoGustavo := new(ContaCorrente)
	// contaDoGustavo.titular = "Gustavo"
	// contaDoGustavo.saldo = 100

	contaDaSilvia := ContaCorrente{titular: "Silvia", saldo: 300}
	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 100}

	fmt.Println(contaDoGustavo)
									// usar & comercial pro ponteiro
	status := contaDaSilvia.Transferir(200,&contaDoGustavo)
	fmt.Println(status)
	fmt.Println("o novo saldo do Gustavo é", contaDoGustavo.saldo)
	fmt.Println("o novo saldo da Silvia é", contaDaSilvia.saldo)


	

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
}															//nao usar new e usar ponteiro
func (c *ContaCorrente) Transferir (valorDaTranferencia float64, contaDestino *ContaCorrente) (bool){
	if valorDaTranferencia < c.saldo && valorDaTranferencia > 0 {
		c.saldo -= valorDaTranferencia
		contaDestino.Depositar(valorDaTranferencia)
		return true

	} else {
		return false
	}
	
}
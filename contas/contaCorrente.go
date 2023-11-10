package contas

import "banco/clientes"

// No terminal, entre na pasta raíz do seu projeto () e digite o comando [ go mod init "NOMEdoPROJETO"]

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

// equivalente a init, self/ 'c' (poderia ser qualquer nome mas em go é a primeira letra do ponteiro) é o endereço de *ContaCorrente // ENtrada   //retorno
func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo += valorDoSaque
		return "Deposito realizado com sucesso! ", c.Saldo
	} else {
		return "Saldo insuficiente! ", c.Saldo
	}
}
func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.Saldo += valorDeposito
		return "Saque realizado com sucesso! seu novo Saldo é ", c.Saldo
	} else {
		return "Saldo insuficiente! seu Saldo é ", c.Saldo
	}
} //nao usar new e usar ponteiro
func (c *ContaCorrente) Transferir(valorDaTranferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTranferencia < c.Saldo && valorDaTranferencia > 0 {
		c.Saldo -= valorDaTranferencia
		contaDestino.Depositar(valorDaTranferencia)
		return true

	} else {
		return false
	}

}

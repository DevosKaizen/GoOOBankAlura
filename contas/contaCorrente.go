package contas

import "banco/clientes"

// No terminal, entre na pasta raíz do seu projeto () e digite o comando [ go mod init "NOMEdoPROJETO"]

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

// equivalente a init, self/ 'c' (poderia ser qualquer nome mas em go é a primeira letra do ponteiro) é o endereço de *ContaCorrente // ENtrada   //retorno
func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo += valorDoSaque
		return "Deposito realizado com sucesso! ", c.saldo
	} else {
		return "saldo insuficiente! ", c.saldo
	}
}
func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Saque realizado com sucesso! seu novo saldo é ", c.saldo
	} else {
		return "saldo insuficiente! seu saldo é ", c.saldo
	}
} //nao usar new e usar ponteiro
func (c *ContaCorrente) Transferir(valorDaTranferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTranferencia < c.saldo && valorDaTranferencia > 0 {
		c.saldo -= valorDaTranferencia
		contaDestino.Depositar(valorDaTranferencia)
		return true

	} else {
		return false
	}

}
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

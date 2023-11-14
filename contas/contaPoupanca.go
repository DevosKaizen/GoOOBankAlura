package contas

import (
	"banco/clientes"
)

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Deposito realizado com sucesso! ", c.saldo
	} else {
		return "saldo insuficiente! ", c.saldo
	}
}
func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Saque realizado com sucesso! seu novo saldo é ", c.saldo
	} else {
		return "saldo insuficiente! seu saldo é ", c.saldo
	}
} //nao usar new e usar ponteiro
func (c *ContaPoupanca) Transferir(valorDaTranferencia float64, contaDestino *ContaPoupanca) bool {
	if valorDaTranferencia < c.saldo && valorDaTranferencia > 0 {
		c.saldo -= valorDaTranferencia
		contaDestino.Depositar(valorDaTranferencia)
		return true

	} else {
		return false
	}

}
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}


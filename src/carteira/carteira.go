package carteira

type Carteira struct {
	Saldo float64
}

func AtualizaSaldo(saldoAtual, valorRecebido float64) float64 {
	return saldoAtual + valorRecebido
}

func Depositar(valor float64) float64 {
	saldoAtual := 45.6
	saldoFinal := valor + saldoAtual
	return saldoFinal
}

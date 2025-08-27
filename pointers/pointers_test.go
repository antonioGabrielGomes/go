package pointers

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		error := carteira.Retirar(Bitcoin(10))

		confirmaSaldo(t, carteira, Bitcoin(10))
		confirmaErroInexistente(t, error)
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)
		confirmaError(t, erro, ErroSaldoInsulficiente.Error())
	})
}

func confirmaSaldo(t *testing.T, carteira Carteira, esperado Bitcoin) {
	t.Helper()
	resultado := carteira.Saldo()

	if resultado != esperado {
		t.Errorf("resultado %s, esperado %s", resultado, esperado)
	}
}

func confirmaError(t *testing.T, valor error, esperado string) {
	t.Helper()
	if valor == nil {
		t.Fatal("esperava um erro, mas nenhum ocorreu")
	}

	resultado := valor.Error()

	if resultado != esperado {
		t.Errorf("resultado %s, esperado %s", resultado, esperado)
	}
}

func confirmaErroInexistente(t *testing.T, resultado error) {
	t.Helper()
	if resultado != nil {
		t.Fatal("erro inesperado recebido")
	}
}

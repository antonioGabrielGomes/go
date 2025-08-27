package pointers

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Carteira struct {
	saldo Bitcoin
}

var ErroSaldoInsulficiente = errors.New("não é possível retirar: saldo insuficiente")

func (carteira *Carteira) Depositar(quantidade Bitcoin) {
	carteira.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade > c.saldo {
		return ErroSaldoInsulficiente
	}
	c.saldo -= quantidade
	return nil
}

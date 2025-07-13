package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Carteira struct {
	saldo Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
	c.saldo += quantidade
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {

	if quantidade > c.saldo {
		return errors.New("Saldo insuficiente")
	}
	c.saldo -= quantidade

	return nil

}

func (c *Carteira) Saldo() Bitcoin {

	return c.saldo
}

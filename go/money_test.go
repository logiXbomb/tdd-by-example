package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	USD = iota
	CHF
)

type Money struct {
	Amount   int
	Currency int
}

func (m *Money) Equal(money *Money) bool {
	return m.Currency == money.Currency && m.Amount == money.Amount
}

func (m *Money) Times(amount int) *Money {
	return NewMoney(m.Amount*amount, m.Currency)
}

func (m *Money) Plus(money *Money) *Money {
	return NewMoney(m.Amount+money.Amount, m.Currency)
}

func NewMoney(amount int, currency int) *Money {
	return &Money{amount, currency}
}

func TestMoney(t *testing.T) {
	t.Run("multiplication", func(t *testing.T) {
		five := NewMoney(5, USD)
		assert.True(t, NewMoney(10, USD).Equal(five.Times(2)))
		assert.True(t, NewMoney(15, USD).Equal(five.Times(3)))
	})

	t.Run("multiplication in Francs", func(t *testing.T) {
		five := NewMoney(5, CHF)
		assert.True(t, NewMoney(10, CHF).Equal(five.Times(2)))
		assert.True(t, NewMoney(15, CHF).Equal(five.Times(3)))
	})

	t.Run("equality", func(t *testing.T) {
		assert.True(t, NewMoney(5, USD).Equal(NewMoney(5, USD)))
		assert.False(t, NewMoney(5, USD).Equal(NewMoney(6, USD)))
		assert.False(t, NewMoney(5, USD).Equal(NewMoney(5, CHF)))
	})

	t.Run("simple addition", func(t *testing.T) {
		sum := NewMoney(5, USD).Plus(NewMoney(5, USD))
		assert.Equal(t, NewMoney(10, USD), sum)
	})
}

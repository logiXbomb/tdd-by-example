package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Dollar struct {
	Amount int
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.Amount * multiplier)
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		Amount: amount,
	}
}

func TestMoney(t *testing.T) {
	t.Run("multiplication", func(t *testing.T) {
		five := NewDollar(5)
		product := five.Times(2)
		assert.Equal(t, 10, product.Amount)
		product = five.Times(3)
		assert.Equal(t, 15, product.Amount)
	})
}

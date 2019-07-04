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

func (d *Dollar) Equal(dollar *Dollar) bool {
	return d.Amount == dollar.Amount
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		Amount: amount,
	}
}

func TestMoney(t *testing.T) {
	t.Run("multiplication", func(t *testing.T) {
		five := NewDollar(5)
		assert.True(t, NewDollar(10).Equal(five.Times(2)))
		assert.True(t, NewDollar(15).Equal(five.Times(3)))
	})

	t.Run("equality", func(t *testing.T) {
		assert.True(t, NewDollar(5).Equal(NewDollar(5)))
		assert.False(t, NewDollar(5).Equal(NewDollar(6)))
	})
}

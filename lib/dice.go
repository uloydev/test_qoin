package lib

import (
	"math/rand"
	"time"
)

type Dice interface {
	Roll()
	GetNumber() int
}

type DiceImpl struct {
	Number int
}

func NewDice() Dice {
	return &DiceImpl{}
}

func (d *DiceImpl) Roll() {
	rand.Seed(time.Now().UnixNano())
	d.Number = rand.Intn(6) + 1
}

func (d *DiceImpl) GetNumber() int {
	return d.Number
}

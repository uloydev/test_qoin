package lib

import (
	"strconv"
	"strings"
)

type Player interface {
	Play()
	RemoveDice(index int)
	AddPoint(point int)
	GetPoint() int
	GetDices() []Dice
	DiceToString() string
	IsFinished() bool
	MergeBonus()
	AddBonus(dice Dice)
}

type PlayerImpl struct {
	Dices     []Dice
	Point     int
	BonusDice []Dice
}

func NewPlayer(diceCount int) Player {
	dices := []Dice{}

	for i := 0; i < diceCount; i++ {
		dices = append(dices, NewDice())
	}

	return &PlayerImpl{
		Dices: dices,
		Point: 0,
	}
}

func (p *PlayerImpl) Play() {
	for _, dice := range p.Dices {
		dice.Roll()
	}
}

func (p *PlayerImpl) RemoveDice(index int) {
	if len(p.Dices) == 1 {
		p.Dices = []Dice{}
	} else if index+1 == len(p.Dices) {
		p.Dices = p.Dices[:index]
	} else if len(p.Dices) != 0 {
		p.Dices = append(p.Dices[:index], p.Dices[index+1:]...)
	}
}

func (p *PlayerImpl) AddPoint(point int) {
	p.Point += point
}

func (p *PlayerImpl) GetPoint() int {
	return p.Point
}

func (p *PlayerImpl) GetDices() []Dice {
	return p.Dices
}

func (p *PlayerImpl) DiceToString() string {
	strArr := []string{}
	for _, dice := range p.Dices {
		strArr = append(strArr, strconv.Itoa(dice.GetNumber()))
	}
	return strings.Join(strArr, ", ")
}

func (p *PlayerImpl) IsFinished() bool {
	return len(p.Dices) == 0 && len(p.BonusDice) == 0
}

func (p *PlayerImpl) MergeBonus() {
	p.Dices = append(p.Dices, p.BonusDice...)
	p.BonusDice = []Dice{}
}

func (p *PlayerImpl) AddBonus(dice Dice) {
	p.BonusDice = append(p.BonusDice, dice)
}

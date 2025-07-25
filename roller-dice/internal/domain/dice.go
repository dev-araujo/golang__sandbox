package domain

type DiceType int

const (
	Dice4Sides  DiceType = 4
	Dice6Sides  DiceType = 6
	Dice8Sides  DiceType = 8
	Dice10Sides DiceType = 10
	Dice20Sides DiceType = 20
)

func (d DiceType) IsValid() bool {
	switch d {
	case Dice4Sides, Dice6Sides, Dice8Sides, Dice10Sides, Dice20Sides:
		return true
	default:
		return false
	}
}

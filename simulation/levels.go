package simulation

import (
	"math"
)

func AbilityModifier(score int) int64 {
	return int64(math.Floor((float64(score) - 10) * .5))
}

func ProficiencyBonus(level int) int64 {
	return int64(math.Ceil(float64(level)*.25) + 1)
}

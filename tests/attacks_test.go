package tests

import (
	"fmt"
	"testing"

	. "example.com/gottrpg/v2/simulation"
)

func TestAttackRoll(t *testing.T) {
	die, i, level, abilityScore := Icosahedron, 3000, 4, 17
	distribution := []int64{}
	for i > 0 {
		roll := DiceRoll(len(die), 1)
		roll += AbilityModifier(abilityScore) + ProficiencyBonus(level)
		distribution = append(distribution, roll)
		i--
	}
	fmt.Println(Mean(distribution))
}

func TestDamage(t *testing.T) {
	die, i, abilityScore := Hexahedron, 3000, 17
	distribution := []int64{}
	for i > 0 {
		roll := DiceRoll(len(die), 1)
		roll += AbilityModifier(abilityScore)
		distribution = append(distribution, roll)
		i--
	}
	fmt.Println(Mean(distribution))
}

// TODO: simulate one or more rounds of combat given character stats and a class or classes
// additionally require at least one enemy with an AC

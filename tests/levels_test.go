package tests

import (
	"testing"

	. "example.com/gottrpg/v2/simulation"
)

func TestAbility(t *testing.T) {
	if AbilityModifier(1) != -5 {
		t.Fail()
	}
	if AbilityModifier(3) != -4 {
		t.Fail()
	}
	if AbilityModifier(30) != 10 {
		t.Fail()
	}
}

func TestProficiency(t *testing.T) {
	if ProficiencyBonus(2) != 2 {
		t.Fail()
	}
	if ProficiencyBonus(5) != 3 {
		t.Fail()
	}
	if ProficiencyBonus(11) != 4 {
		t.Fail()
	}
	if ProficiencyBonus(16) != 5 {
		t.Fail()
	}
	if ProficiencyBonus(19) != 6 {
		t.Fail()
	}
}

package character

import (
	sim "example.com/gottrpg/v2/simulation"
)

type Spell struct {
	CastingTime int
	Components  []Component
	Duration    int
	Effect      Effect
	Name        string
	Range       int64 // if 0 (self)
	School      School
	Target      Target
}

type Effect struct {
	Die      sim.Die
	Number   int
	Modifier []Amplification
}

type Amplification struct {
	Level int // level amplification occurs at
	Dice  int // number of dice to add
}

type Target string

var (
	SELF     Target = "self"
	FRIENDLY Target = "friendly"
	HOSTILE  Target = "hostile"
)

type School string

var (
	ABJURATION    School = "abjuration"
	CONJURATION   School = "conjuration"
	DIVINATION    School = "divination"
	ENCHANTMENT   School = "enchantment"
	EVOCATION     School = "evocation"
	ILLUSION      School = "illusion"
	NECROMANCY    School = "necromancy"
	TRANSMUTATION School = "transmutation"
)

type Component string

var (
	MATERIAL Component = "material"
	SOMATIC  Component = "somatic"
	VERBAL   Component = "verbal"
)

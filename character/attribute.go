package character

type Attr string

const Strength Attr = "STR"
const Dexterity Attr = "DEX"
const Constitution Attr = "CON"
const Intelligence Attr = "INT"
const Wisdom Attr = "WIS"
const Charisma Attr = "CHA"

type Attribute struct {
	Name  Attr
	Value int
}

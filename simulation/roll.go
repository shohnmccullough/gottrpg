package simulation

// D4 rolls a 4-sided die (tetrahedron) the specified number of times
func D4(quantity int) int64 {
	return DiceRoll(4, quantity)
}

// D6 rolls a 6-sided die (hexahedron) the specified number of times
func D6(quantity int) int64 {
	return DiceRoll(6, quantity)
}

// D8 rolls an 8-sided die (octahedron) the specified number of times
func D8(quantity int) int64 {
	return DiceRoll(8, quantity)
}

// D10 rolls a 10-sided die (decagon) the specified number of times
func D10(quantity int) int64 {
	return DiceRoll(10, quantity)
}

// D12 rolls a 12-sided die (dodecahedron) the specified number of times
func D12(quantity int) int64 {
	return DiceRoll(12, quantity)
}

// D20 rolls a 20-sided die (icosahedron) the specified number of times
func D20(quantity int) int64 {
	return DiceRoll(20, quantity)
}
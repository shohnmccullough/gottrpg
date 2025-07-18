package main

import (
	"fmt"

	char "example.com/gottrpg/v2/character"
	sim "example.com/gottrpg/v2/simulation"
)

// CreateSampleCharacter creates a sample D&D 5e character (Elara, Half-Elf Ranger)
func CreateSampleCharacter() char.Character {
	// Character attributes (standard array with racial bonuses)
	attributes := []char.Attribute{
		{Name: char.Strength, Value: 13},
		{Name: char.Dexterity, Value: 16}, // 15 + 1 (Half-Elf)
		{Name: char.Constitution, Value: 14},
		{Name: char.Intelligence, Value: 10},
		{Name: char.Wisdom, Value: 15},
		{Name: char.Charisma, Value: 13}, // 12 + 1 (Half-Elf)
	}

	// Half-Elf race
	halfElfRace := char.Race{
		Name:                "Half-Elf",
		AbilityScoreBonuses: []int{0, 1, 0, 0, 0, 2}, // +1 DEX, +2 CHA
		Size:                5,                       // Medium
		Speed:               30,
		Vision:              "Darkvision 60 feet",
		Languages:           []string{"Common", "Elvish", "Orcish"},
		Alignment: char.Alignment{
			Morality:         "Good",
			SocietalAttitude: "Chaotic",
		},
		Subrace: char.Subrace{
			Name: "Standard Half-Elf",
		},
	}

	// Ranger class
	rangerClass := char.Class{
		Name:     "Ranger",
		Level:    3,
		Primary:  true,
		Features: []char.Feature{}, // TODO: Add ranger features
	}

	// Folk Hero background
	background := char.Background{
		Name: "Folk Hero",
		Skills: []char.Skill{
			"Animal Handling",
			"Survival",
		},
		Languages: []string{"Giant"},
		Tools:     []string{"Smith's Tools", "Vehicles (Land)"},
		Gold:      10,
	}

	// Proficiencies (Ranger + Background + Racial)
	// proficiencies := []*char.Proficiency{
	// 	char.AnimalHandling.Proficient(),
	// 	char.Athletics.Proficient(),
	// 	char.Insight.Proficient(),
	// 	char.Investigation.Proficient(),
	// 	char.Nature.Proficient(),
	// 	char.Perception.Proficient(),
	// 	char.Stealth.Proficient(),
	// 	char.Survival.Proficient(),
	// }

	// Saving throws (Rangers are proficient in STR and DEX)
	savingThrows := map[string]int{
		"STR": 3, // 1 (mod) + 2 (prof bonus)
		"DEX": 5, // 3 (mod) + 2 (prof bonus)
		"CON": 2, // 2 (mod)
		"INT": 0, // 0 (mod)
		"WIS": 2, // 2 (mod)
		"CHA": 1, // 1 (mod)
	}

	// Sample ranger spells
	spells := []char.Spell{
		{
			Name:        "Hunter's Mark",
			School:      char.DIVINATION,
			CastingTime: 1, // 1 bonus action
			Range:       90,
			Duration:    3600, // 1 hour in seconds
			Components:  []char.Component{char.VERBAL},
			Target:      char.HOSTILE,
			Effect: char.Effect{
				Die:      sim.Decagon,
				Number:   1,
				Modifier: []char.Amplification{},
			},
		},
		{
			Name:        "Cure Wounds",
			School:      char.EVOCATION,
			CastingTime: 1, // 1 action
			Range:       0, // Touch
			Duration:    0, // Instantaneous
			Components:  []char.Component{char.VERBAL, char.SOMATIC},
			Target:      char.FRIENDLY,
			Effect: char.Effect{
				Die:    sim.Octahedron,
				Number: 1,
				Modifier: []char.Amplification{
					{Level: 2, Dice: 1},
					{Level: 3, Dice: 2},
				},
			},
		},
	}

	return char.Character{
		Name:       "Elara Moonwhisper",
		Level:      3,
		Experience: 900,
		Background: background,
		Race:       halfElfRace,
		Class:      []char.Class{rangerClass},
		Attributes: attributes,
		// Proficiencies:    proficiencies,
		SavingThrows:     savingThrows,
		ProficiencyBonus: 2,
		Feat:             []string{}, // No feats at level 3
		Spells:           spells,
	}
}

// DisplayCharacter prints character information
func DisplayCharacter(c char.Character) {
	fmt.Printf("=== %s ===\n", c.Name)
	fmt.Printf("Level %d %s %s\n", c.Level, c.Race.Name, c.Class[0].Name)
	fmt.Printf("Background: %s\n", c.Background.Name)
	fmt.Printf("Experience: %d\n", c.Experience)
	fmt.Printf("Alignment: %s %s\n", c.Race.Alignment.SocietalAttitude, c.Race.Alignment.Morality)

	fmt.Println("\n--- Attributes ---")
	for _, attr := range c.Attributes {
		modifier := (attr.Value - 10) / 2
		fmt.Printf("%s: %d (%+d)\n", attr.Name, attr.Value, modifier)
	}

	fmt.Printf("\nProficiency Bonus: +%d\n", c.ProficiencyBonus)

	fmt.Println("\n--- Saving Throws ---")
	for attr, bonus := range c.SavingThrows {
		fmt.Printf("%s: %+d\n", attr, bonus)
	}

	fmt.Println("\n--- Skills ---")
	for _, prof := range c.Proficiencies {
		if prof.Proficiency {
			expertiseText := ""
			if prof.Expertise {
				expertiseText = " (Expertise)"
			}
			fmt.Printf("%s%s\n", prof.Name, expertiseText)
		}
	}

	fmt.Println("\n--- Spells ---")
	for _, spell := range c.Spells {
		fmt.Printf("%s (%s)\n", spell.Name, spell.School)
	}

	fmt.Printf("\n--- Languages ---")
	allLanguages := append(c.Race.Languages, c.Background.Languages...)
	for i, lang := range allLanguages {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(lang)
	}
	fmt.Println()
}

// func main() {
// 	character := CreateSampleCharacter()
// 	DisplayCharacter(character)
// }

package character

// TODO: revisit this

/* Create a new Proficiency struct pointer */
func ProficiencyBuilder(n Skill, a Attr) *Proficiency {
	return &Proficiency{
		Name:        n,
		Attribute:   a,
		Proficiency: false,
		Expertise:   false,
	}
}

type Proficiency struct {
	Name        Skill
	Attribute   Attr
	Proficiency bool
	Expertise   bool
}

/* Set Proficiency to true */
func (p *Proficiency) Proficient() *Proficiency {
	p.Proficiency = true
	return p
}

/* Set Expertise to true */
func (p *Proficiency) Expert() *Proficiency {
	p.Expertise = true
	return p
}

/* Strength-based skill proficiencies */

var Athletics *Proficiency = ProficiencyBuilder(athletics, Strength)

/* Dexterity-based skill proficiencies */

var Acrobatics *Proficiency = ProficiencyBuilder(acrobatics, Dexterity)

var SleightOfHand *Proficiency = ProficiencyBuilder(sleightOfHand, Dexterity)

var Stealth *Proficiency = ProficiencyBuilder(stealth, Dexterity)

/* Intelligence-based skill proficiencies */

var Arcana *Proficiency = ProficiencyBuilder(arcana, Intelligence)

var History *Proficiency = ProficiencyBuilder(history, Intelligence)

var Investigation *Proficiency = ProficiencyBuilder(investigation, Intelligence)

var Nature *Proficiency = ProficiencyBuilder(nature, Intelligence)

var Religion *Proficiency = ProficiencyBuilder(religion, Intelligence)

/* Wisdom-based skill proficiencies */

var AnimalHandling *Proficiency = ProficiencyBuilder(animalHandling, Wisdom)

var Insight *Proficiency = ProficiencyBuilder(insight, Wisdom)

var Medicine *Proficiency = ProficiencyBuilder(medicine, Wisdom)

var Perception *Proficiency = ProficiencyBuilder(perception, Wisdom)

var Survival *Proficiency = ProficiencyBuilder(survivor, Wisdom)

/* Charisma-based skill proficiencies */

var Deception *Proficiency = ProficiencyBuilder(deception, Charisma)

var Intimidation *Proficiency = ProficiencyBuilder(intimidation, Charisma)

var Performance *Proficiency = ProficiencyBuilder(performance, Charisma)

var Persuasion *Proficiency = ProficiencyBuilder(persuasion, Charisma)

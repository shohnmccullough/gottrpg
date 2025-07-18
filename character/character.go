package character

type Character struct {
	Name             string
	Level            int
	Experience       int64
	Background       Background
	Race             Race
	Class            []Class // allow slice for multi-classing
	Attributes       []Attribute
	Proficiencies    []Proficiency
	SavingThrows     map[string]int // should be a struct to track proficiency
	ProficiencyBonus int
	Feat             []string
	Spells           []Spell
}

/* ----------------------------------------------------------------------------------------- */

type Background struct {
	Name      string
	Skills    []Skill
	Languages []string // should share "enums" with Race(s)
	Tools     []string // should share "enums"
	Gold      int      // TODO: somehow handle this sensically, seems like there needs to be a running character sheet and an underlying struct it starts from
}

type Class struct {
	Name     string
	Level    int
	Features []Feature
	Primary  bool // represent primary class (i.e. the initial class of a character that has multi-classed)
}

type Feature struct{}

/* ----------------------------------------------------------------------------------------- */

type Race struct {
	AbilityScoreBonuses []int
	Alignment           Alignment
	Languages           []string
	Name                string
	Size                int
	Speed               int
	Vision              string // TODO: struct to handle this
	Subrace             Subrace
}

type Subrace struct {
	Name string
}

type Morality string
type SocietalAttitude string

const (
	Evil      Morality         = "Evil"
	Good      Morality         = "Good"
	NeutralM  Morality         = "Neutral"
	Chaotic   SocietalAttitude = "Chaotic"
	Lawful    SocietalAttitude = "Lawful"
	NeutralSA SocietalAttitude = "Neutral"
)

type Alignment struct {
	Morality         string
	SocietalAttitude string
}

// TODO: how are bonuses/enhancements/feats/etc. applied?
/**

types of actions
----------------
Saving Throw
Skill Check
Spell / Attack

*/

type ActionModifier struct {
	Effect int
}

package simulation

type Condition int

/** RAW conditions */
const (
	Blinded       = 0
	Charmed       = 0
	Deafened      = 0
	Exhaustion    = 0 // TODO: need to represent this in multiple levels
	Frightened    = 0
	Grappled      = 0
	Incapacitated = 0
	Invisible     = 0
	Paralyzed     = 0
	Petrified     = 0
	Poisoned      = 0
	Prone         = 0
	Restrained    = 0
	Stunned       = 0
	Unconscious   = 0
)

/** Condition-like effects within the RAW */
const (
	Acid      = 0
	Burning   = 0
	Chilled   = 0
	Enwebbed  = 0
	Hasted    = 0
	Slowed    = 0
	Surprised = 0
)

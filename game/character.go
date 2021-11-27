package game

import "math"

// Character -
type Character struct {
	Character *CharacterData `json:"character"`
	MaxHP     int            `json:"maxhp"`
	HP        int            `json:"hp"`
	Attack    int            `json:"attack"`
	Defence   int            `json:"defence"`
}

// NewCharacter -new a Character with CharacterData
func NewCharacter(cd *CharacterData) *Character {
	return &Character{
		Character: cd,
		MaxHP:     cd.HP,
		HP:        cd.HP,
		Attack:    cd.Attack,
		Defence:   cd.Defence,
	}
}

// CanAttack -
func (c *Character) CanAttack(c1 *Character) bool {
	return c.Attack > c1.Defence
}

// StartAttack - return isKO
func (c *Character) StartAttack(c1 *Character) bool {
	if !c.CanAttack(c1) {
		return false
	}

	c1.HP -= c.Attack - c1.Defence

	if c1.HP <= 0 {
		c1.HP = 0

		return true
	}

	return false
}

// CalcMinDataWithAttackTurns - minact, maxhp
func (c *Character) CalcMinDataWithAttackTurns(params *MinDataWithAttackTurnsParams) *CharacterDataArea {
	if !params.IsValid() {
		return nil
	}

	turnwin := math.Ceil(float64(c.Character.HP) / float64(params.Turns))
	minact := int(turnwin) + c.Defence

	maxhp := math.Ceil(float64(c.Attack*(params.Turns-1)) / params.MaxLostHPPer)

	mindef := math.Ceil(float64(c.Character.Attack) * params.MinDefPer)
	maxdef := math.Ceil(float64(c.Character.Attack) * params.MaxDefPer)

	mindefhp := math.Ceil(float64((c.Attack-int(mindef))*(params.Turns-1)) / params.MaxLostHPPer)
	maxdefhp := math.Ceil(float64((c.Attack-int(maxdef))*(params.Turns-1)) / params.MaxLostHPPer)

	return &CharacterDataArea{
		MinAttack:    minact,
		NoDefenceHP:  int(maxhp),
		MinDef:       int(mindef),
		MaxDef:       int(maxdef),
		MinDefenceHP: int(mindefhp),
		MaxDefenceHP: int(maxdefhp),
	}
}

// Clone - Clone a Character
func (c *Character) Clone() *Character {
	return NewCharacter(c.Character)
}

// Team -
type Team struct {
	Characters []*Character `json:"characters"`
}

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
func (c *Character) CalcMinDataWithAttackTurns(turns int, maxLostHPPer float64, minDefPer float64, maxDefPer float64) *CharacterDataArea {
	if turns <= 0 || maxLostHPPer >= 1 {
		return nil
	}

	turnwin := math.Ceil(float64(c.Character.HP) / float64(turns))
	minact := int(turnwin) + c.Defence

	maxhp := math.Ceil(float64(c.Attack*(turns-1)) / maxLostHPPer)

	mindef := math.Ceil(float64(c.Character.Attack) * minDefPer)
	maxdef := math.Ceil(float64(c.Character.Attack) * maxDefPer)

	mindefhp := math.Ceil(float64((c.Attack-int(mindef))*(turns-1)) / maxLostHPPer)
	maxdefhp := math.Ceil(float64((c.Attack-int(maxdef))*(turns-1)) / maxLostHPPer)

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

package game

// CharacterData -
type CharacterData struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	HP     int    `json:"hp"`
	Attack int    `json:"attack"`
}

// Clone - Clone a CharacterData
func (cd *CharacterData) Clone() *CharacterData {
	return &CharacterData{
		ID:     cd.ID,
		Name:   cd.Name,
		HP:     cd.HP,
		Attack: cd.Attack,
	}
}

// CharacterDataIndex -
type CharacterDataIndex struct {
	IDIndex     int
	NameIndex   int
	HPIndex     int
	AttackIndex int
}

// Character -
type Character struct {
	Character *CharacterData `json:"character"`
	MaxHP     int            `json:"maxhp"`
	HP        int            `json:"hp"`
	Attack    int            `json:"attack"`
}

// NewCharacter -new a Character with CharacterData
func NewCharacter(cd *CharacterData) *Character {
	return &Character{
		Character: cd,
		MaxHP:     cd.HP,
		HP:        cd.HP,
		Attack:    cd.Attack,
	}
}

// StartAttack - return isKO
func (c *Character) StartAttack(c1 *Character) bool {
	c1.HP -= c.Attack

	if c1.HP <= 0 {
		c1.HP = 0

		return true
	}

	return false
}

// Team -
type Team struct {
	Characters []*Character `json:"characters"`
}

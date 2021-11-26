package game

// CharacterData -
type CharacterData struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	HP      int    `json:"hp"`
	Attack  int    `json:"attack"`
	Defence int    `json:"defence"`
	Gold    int    `json:"gold"`
	Exp     int    `json:"exp"`
}

// Clone - Clone a CharacterData
func (cd *CharacterData) Clone() *CharacterData {
	return &CharacterData{
		ID:      cd.ID,
		Name:    cd.Name,
		HP:      cd.HP,
		Attack:  cd.Attack,
		Defence: cd.Defence,
		Gold:    cd.Gold,
		Exp:     cd.Exp,
	}
}

// CharacterDataIndex -
type CharacterDataIndex struct {
	IDIndex      int
	NameIndex    int
	HPIndex      int
	AttackIndex  int
	DefenceIndex int
	GoldIndex    int
	ExpIndex     int
}

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

// Clone - Clone a Character
func (c *Character) Clone() *Character {
	return NewCharacter(c.Character)
}

// Team -
type Team struct {
	Characters []*Character `json:"characters"`
}

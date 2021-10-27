package game

// CharacterData -
type CharacterData struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	HP     int    `json:"hp"`
	Attack int    `json:"attack"`
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
	Character CharacterData `json:"character"`
}

// Team -
type Team struct {
	Characters []*Character `json:"characters"`
}

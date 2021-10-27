package game

// CharacterData -
type CharacterData struct {
	HP     int `json:"hp"`
	Attack int `json:"attack"`
}

// Character -
type Character struct {
	Character CharacterData `json:"character"`
}

// Team -
type Team struct {
	Characters []*Character `json:"characters"`
}

package game

// Character -
type Character struct {
	HP     int `json:"hp"`
	Attack int `json:"attack"`
}

// Team -
type Team struct {
	Characters []*Character `json:"characters"`
}

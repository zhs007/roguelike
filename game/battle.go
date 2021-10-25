package game

// Battle -
type Battle struct {
	Teams []*Team `json:"teams"`
}

// OnBattle -
func (battle *Battle) OnBattle() error {
	return nil
}

package game

// Battle -
type Battle struct {
	Characters []*Character `json:"characters"`
}

// NewBattle - new a battle
func NewBattle(c0 *Character, c1 *Character) *Battle {
	return &Battle{
		Characters: []*Character{c0, c1},
	}
}

// OnBattle -
func (battle *Battle) OnBattle() int {
	for {
		c0ko := battle.Characters[0].StartAttack(battle.Characters[1])
		if c0ko {
			return 1
		}

		c1ko := battle.Characters[1].StartAttack(battle.Characters[0])
		if c1ko {
			return -1
		}
	}
}

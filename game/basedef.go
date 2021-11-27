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

// CharacterDataArea -
type CharacterDataArea struct {
	NoDefenceHP  int `json:"noDefenceHP"`
	MinAttack    int `json:"minAttack"`
	MinDef       int `json:"minDefence"`
	MaxDef       int `json:"maxDefence"`
	MinDefenceHP int `json:"minDefenceHP"`
	MaxDefenceHP int `json:"maxDefenceHP"`
}

// MinDataWithAttackTurnsParams - for CalcMinDataWithAttackTurns
type MinDataWithAttackTurnsParams struct {
	Turns        int
	MaxLostHPPer float64
	MinDefPer    float64
	MaxDefPer    float64
}

func (params *MinDataWithAttackTurnsParams) IsValid() bool {
	if params.Turns <= 0 || params.MaxLostHPPer >= 1 {
		return false
	}

	return true
}

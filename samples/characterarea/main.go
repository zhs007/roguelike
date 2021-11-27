package main

import (
	"github.com/zhs007/goutils"
	"github.com/zhs007/roguelike"
	"github.com/zhs007/roguelike/game"
	"go.uber.org/zap"
)

func main() {
	goutils.InitLogger("characterarea", roguelike.Version, "debug", true, "./")

	mgrCharacter, err := game.LoadCharacterMgr("./gamedata/character.xlsx")
	if err != nil {
		goutils.Error("LoadCharacterMgr",
			zap.Error(err))

		return
	}

	character, err := mgrCharacter.NewCharacter(1)
	if err != nil {
		goutils.Error("NewCharacter",
			zap.Int("cid", 1),
			zap.Error(err))

		return
	}

	cda := character.CalcMinDataWithAttackTurns(2, 0.2, 0.4, 0.7)
	goutils.Info("CalcMinDataWithAttackTurns",
		goutils.JSON("character", character),
		goutils.JSON("CharacterDataArea", cda))
}

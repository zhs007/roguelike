package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
	"github.com/zhs007/goutils"
	"github.com/zhs007/roguelike"
	"github.com/zhs007/roguelike/game"
	"go.uber.org/zap"
)

func saveCharacterDataArea(mgrCharacter *game.CharacterMgr, fn string, params []*game.MinDataWithAttackTurnsParams) {
	f := excelize.NewFile()

	sheet := "Sheet1"

	f.SetCellValue(sheet, goutils.Pos2Cell(0, 0), "cid")
	f.SetCellValue(sheet, goutils.Pos2Cell(1, 0), "name")
	f.SetCellValue(sheet, goutils.Pos2Cell(2, 0), "hp")
	f.SetCellValue(sheet, goutils.Pos2Cell(3, 0), "attack")
	f.SetCellValue(sheet, goutils.Pos2Cell(4, 0), "defence")
	f.SetCellValue(sheet, goutils.Pos2Cell(5, 0), "gold")
	f.SetCellValue(sheet, goutils.Pos2Cell(6, 0), "exp")

	minx := 7
	cdalen := 6

	for i := range params {
		f.SetCellValue(sheet, goutils.Pos2Cell(minx+i*cdalen, 0), fmt.Sprintf("noDefenceHP%v", i))
		f.SetCellValue(sheet, goutils.Pos2Cell(minx+i*cdalen+1, 0), fmt.Sprintf("minAttack%v", i))
		f.SetCellValue(sheet, goutils.Pos2Cell(minx+i*cdalen+2, 0), fmt.Sprintf("minDefence%v", i))
		f.SetCellValue(sheet, goutils.Pos2Cell(minx+i*cdalen+3, 0), fmt.Sprintf("maxDefence%v", i))
		f.SetCellValue(sheet, goutils.Pos2Cell(minx+i*cdalen+4, 0), fmt.Sprintf("minDefenceHP%v", i))
		f.SetCellValue(sheet, goutils.Pos2Cell(minx+i*cdalen+5, 0), fmt.Sprintf("maxDefenceHP%v", i))
	}

	minx2 := minx + cdalen*len(params)
	f.SetCellValue(sheet, goutils.Pos2Cell(minx2, 0), "minHP")
	f.SetCellValue(sheet, goutils.Pos2Cell(minx2+1, 0), "maxHP")
	f.SetCellValue(sheet, goutils.Pos2Cell(minx2+2, 0), "minAttack")
	f.SetCellValue(sheet, goutils.Pos2Cell(minx2+3, 0), "maxAttack")
	f.SetCellValue(sheet, goutils.Pos2Cell(minx2+4, 0), "minDefence")
	f.SetCellValue(sheet, goutils.Pos2Cell(minx2+5, 0), "maxDefence")

	row := 1
	mgrCharacter.ForEachCharacterData(func(cd *game.CharacterData) {
		f.SetCellValue(sheet, goutils.Pos2Cell(0, row), cd.ID)
		f.SetCellValue(sheet, goutils.Pos2Cell(1, row), cd.Name)
		f.SetCellValue(sheet, goutils.Pos2Cell(2, row), cd.HP)
		f.SetCellValue(sheet, goutils.Pos2Cell(3, row), cd.Attack)
		f.SetCellValue(sheet, goutils.Pos2Cell(4, row), cd.Defence)
		f.SetCellValue(sheet, goutils.Pos2Cell(5, row), cd.Gold)
		f.SetCellValue(sheet, goutils.Pos2Cell(6, row), cd.Exp)

		minhp := 0
		maxhp := 0
		minack := 0
		maxack := 0
		mindef := 0
		maxdef := 0

		for i, v := range params {
			c := game.NewCharacter(cd)
			cda := c.CalcMinDataWithAttackTurns(v)

			f.SetCellValue(sheet, goutils.Pos2Cell(minx+i*cdalen, row), fmt.Sprintf("%v", cda.NoDefenceHP))
			f.SetCellValue(sheet, goutils.Pos2Cell(minx+i*cdalen+1, row), fmt.Sprintf("%v", cda.MinAttack))
			f.SetCellValue(sheet, goutils.Pos2Cell(minx+i*cdalen+2, row), fmt.Sprintf("%v", cda.MinDef))
			f.SetCellValue(sheet, goutils.Pos2Cell(minx+i*cdalen+3, row), fmt.Sprintf("%v", cda.MaxDef))
			f.SetCellValue(sheet, goutils.Pos2Cell(minx+i*cdalen+4, row), fmt.Sprintf("%v", cda.MinDefenceHP))
			f.SetCellValue(sheet, goutils.Pos2Cell(minx+i*cdalen+5, row), fmt.Sprintf("%v", cda.MaxDefenceHP))

			if i == 0 {
				minhp = cda.NoDefenceHP
				maxhp = cda.NoDefenceHP
				minack = cda.MinAttack
				maxack = cda.MinAttack
				mindef = cda.MinDef
				maxdef = cda.MaxDef

				if minhp > cda.MinDefenceHP {
					minhp = cda.MinDefenceHP
				}

				if maxhp < cda.MinDefenceHP {
					maxhp = cda.MinDefenceHP
				}

				if minhp > cda.MaxDefenceHP {
					minhp = cda.MaxDefenceHP
				}

				if maxhp < cda.MaxDefenceHP {
					maxhp = cda.MaxDefenceHP
				}
			} else {
				if minhp > cda.NoDefenceHP {
					minhp = cda.NoDefenceHP
				}

				if maxhp < cda.NoDefenceHP {
					maxhp = cda.NoDefenceHP
				}

				if minhp > cda.MinDefenceHP {
					minhp = cda.MinDefenceHP
				}

				if maxhp < cda.MinDefenceHP {
					maxhp = cda.MinDefenceHP
				}

				if minhp > cda.MaxDefenceHP {
					minhp = cda.MaxDefenceHP
				}

				if maxhp < cda.MaxDefenceHP {
					maxhp = cda.MaxDefenceHP
				}

				if minack > cda.MinAttack {
					minack = cda.MinAttack
				}

				if maxack < cda.MinAttack {
					maxack = cda.MinAttack
				}

				if mindef > cda.MinDef {
					mindef = cda.MinDef
				}

				if maxdef < cda.MaxDef {
					maxdef = cda.MaxDef
				}
			}
		}

		f.SetCellValue(sheet, goutils.Pos2Cell(minx2, row), fmt.Sprintf("%v", minhp))
		f.SetCellValue(sheet, goutils.Pos2Cell(minx2+1, row), fmt.Sprintf("%v", maxhp))
		f.SetCellValue(sheet, goutils.Pos2Cell(minx2+2, row), fmt.Sprintf("%v", minack))
		f.SetCellValue(sheet, goutils.Pos2Cell(minx2+3, row), fmt.Sprintf("%v", maxack))
		f.SetCellValue(sheet, goutils.Pos2Cell(minx2+4, row), fmt.Sprintf("%v", mindef))
		f.SetCellValue(sheet, goutils.Pos2Cell(minx2+5, row), fmt.Sprintf("%v", maxdef))

		row++
	})

	f.SaveAs(fn)
}

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

	cda := character.CalcMinDataWithAttackTurns(&game.MinDataWithAttackTurnsParams{
		Turns:        2,
		MaxLostHPPer: 0.2,
		MinDefPer:    0.4,
		MaxDefPer:    0.8,
	})
	goutils.Info("CalcMinDataWithAttackTurns",
		goutils.JSON("character", character),
		goutils.JSON("CharacterDataArea", cda))

	saveCharacterDataArea(mgrCharacter, "characterarea.xlsx", []*game.MinDataWithAttackTurnsParams{
		{
			Turns:        2,
			MaxLostHPPer: 0.2,
			MinDefPer:    0.4,
			MaxDefPer:    0.8,
		},
		{
			Turns:        3,
			MaxLostHPPer: 0.2,
			MinDefPer:    0.4,
			MaxDefPer:    0.8,
		},
		{
			Turns:        4,
			MaxLostHPPer: 0.2,
			MinDefPer:    0.4,
			MaxDefPer:    0.8,
		},
		{
			Turns:        5,
			MaxLostHPPer: 0.2,
			MinDefPer:    0.4,
			MaxDefPer:    0.8,
		},
	})
}

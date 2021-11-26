package game

import (
	"strings"

	"github.com/xuri/excelize/v2"
	"github.com/zhs007/goutils"
	"go.uber.org/zap"
)

type CharacterMgr struct {
	MapCharacters map[int]*CharacterData
}

func LoadCharacterMgr(fn string) (*CharacterMgr, error) {
	f, err := excelize.OpenFile(fn)
	if err != nil {
		goutils.Error("LoadCharacterMgr:OpenFile",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	lstname := f.GetSheetList()
	if len(lstname) <= 0 {
		goutils.Error("LoadCharacterMgr:GetSheetList",
			goutils.JSON("SheetList", lstname),
			zap.String("fn", fn),
			zap.Error(ErrInvalidCharacterFiles))

		return nil, ErrInvalidCharacterFiles
	}

	rows, err := f.GetRows(lstname[0])
	if err != nil {
		goutils.Error("LoadCharacterMgr:GetRows",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	mgr := &CharacterMgr{
		MapCharacters: make(map[int]*CharacterData),
	}

	cdi := CharacterDataIndex{}

	for y, row := range rows {
		cd := &CharacterData{}

		if y == 0 {
			for x, colCell := range row {
				c := strings.ToLower(colCell)
				if c == "cid" {
					cdi.IDIndex = x
				} else if c == "name" {
					cdi.NameIndex = x
				} else if c == "hp" {
					cdi.HPIndex = x
				} else if c == "attack" {
					cdi.AttackIndex = x
				} else if c == "defence" {
					cdi.DefenceIndex = x
				} else if c == "gold" {
					cdi.GoldIndex = x
				} else if c == "exp" {
					cdi.ExpIndex = x
				}
			}
		} else {
			for x, colCell := range row {

				if x == cdi.IDIndex {
					cid, err := goutils.String2Int64(colCell)
					if err != nil {
						goutils.Error("LoadCharacterMgr:id",
							zap.String("id", colCell),
							zap.Error(err))

						return nil, err
					}

					cd.ID = int(cid)
				} else if x == cdi.NameIndex {
					cd.Name = strings.TrimSpace(colCell)
				} else if x == cdi.HPIndex {
					hp, err := goutils.String2Int64(colCell)
					if err != nil {
						goutils.Error("LoadCharacterMgr:hp",
							zap.String("hp", colCell),
							zap.Error(err))

						return nil, err
					}

					cd.HP = int(hp)
				} else if x == cdi.AttackIndex {
					attack, err := goutils.String2Int64(colCell)
					if err != nil {
						goutils.Error("LoadCharacterMgr:attack",
							zap.String("attack", colCell),
							zap.Error(err))

						return nil, err
					}

					cd.Attack = int(attack)
				} else if x == cdi.DefenceIndex {
					defence, err := goutils.String2Int64(colCell)
					if err != nil {
						goutils.Error("LoadCharacterMgr:defence",
							zap.String("defence", colCell),
							zap.Error(err))

						return nil, err
					}

					cd.Defence = int(defence)
				} else if x == cdi.GoldIndex {
					gold, err := goutils.String2Int64(colCell)
					if err != nil {
						goutils.Error("LoadCharacterMgr:gold",
							zap.String("gold", colCell),
							zap.Error(err))

						return nil, err
					}

					cd.Gold = int(gold)
				} else if x == cdi.ExpIndex {
					exp, err := goutils.String2Int64(colCell)
					if err != nil {
						goutils.Error("LoadCharacterMgr:exp",
							zap.String("exp", colCell),
							zap.Error(err))

						return nil, err
					}

					cd.Exp = int(exp)
				}
			}

			mgr.MapCharacters[cd.ID] = cd
		}
	}

	return mgr, nil
}

func (mgr *CharacterMgr) NewCharacter(cid int) (*CharacterData, error) {
	cd, isok := mgr.MapCharacters[cid]
	if isok {
		return cd, nil
	}

	goutils.Warn("CharacterMgr.NewCharacter",
		zap.Int("cid", cid),
		zap.Error(ErrInvalidCharacterID))

	return nil, ErrInvalidCharacterID
}

package game

import (
	"strings"

	"github.com/xuri/excelize/v2"
	"github.com/zhs007/goutils"
	"go.uber.org/zap"
)

type CharacterMgr struct {
	Characters []*CharacterData
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

	mgr := &CharacterMgr{}
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
					cid, err := goutils.String2Int64(colCell)
					if err != nil {
						goutils.Error("LoadCharacterMgr:hp",
							zap.String("hp", colCell),
							zap.Error(err))

						return nil, err
					}

					cd.HP = int(cid)
				} else if x == cdi.AttackIndex {
					cid, err := goutils.String2Int64(colCell)
					if err != nil {
						goutils.Error("LoadCharacterMgr:attack",
							zap.String("attack", colCell),
							zap.Error(err))

						return nil, err
					}

					cd.Attack = int(cid)
				}
			}

			mgr.Characters = append(mgr.Characters, cd)
		}
	}

	return mgr, nil
}

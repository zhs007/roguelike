package game

import (
	"fmt"

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

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

	return nil, nil
}

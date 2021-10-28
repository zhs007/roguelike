package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
	"github.com/zhs007/goutils"
	"github.com/zhs007/roguelike/game"
)

func main() {
	f := excelize.NewFile()
	tv := 100
	for hp := 1; hp < tv; hp++ {
		cd0 := &game.CharacterData{
			HP:     hp,
			Attack: tv - hp,
		}
		c0 := game.NewCharacter(cd0)

		f.SetCellValue("Sheet1", goutils.Pos2Cell(hp, 0), fmt.Sprintf("hp%vatk%v", hp, tv-hp))

		for hp1 := 1; hp1 < tv; hp1++ {
			f.SetCellValue("Sheet1", goutils.Pos2Cell(0, hp1), fmt.Sprintf("hp%vatk%v", hp1, tv-hp1))
			// if hp1 == hp {
			// 	break
			// }

			cd1 := &game.CharacterData{
				HP:     hp1,
				Attack: tv - hp1,
			}
			c1 := game.NewCharacter(cd1)

			b := game.NewBattle(c0, c1)
			ret := b.OnBattle()
			// fmt.Printf("%v", ret)
			f.SetCellValue("Sheet1", goutils.Pos2Cell(hp, hp1), fmt.Sprintf("%v", ret))
		}
	}

	f.SaveAs("combatpower.xlsx")
}

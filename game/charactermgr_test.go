package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoadCharacterMgr(t *testing.T) {
	mgr, err := LoadCharacterMgr("../gamedata/character.xlsx")
	assert.NoError(t, err)
	assert.Equal(t, len(mgr.Characters), 2)

	assert.Equal(t, mgr.Characters[0].ID, 1)
	assert.Equal(t, mgr.Characters[0].Name, "主角")
	assert.Equal(t, mgr.Characters[0].HP, 100)
	assert.Equal(t, mgr.Characters[0].Attack, 30)

	t.Logf("Test_LoadCharacterMgr OK")
}

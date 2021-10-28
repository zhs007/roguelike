package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoadCharacterMgr(t *testing.T) {
	mgr, err := LoadCharacterMgr("../gamedata/character.xlsx")
	assert.NoError(t, err)
	assert.Equal(t, len(mgr.MapCharacters), 2)

	assert.Equal(t, mgr.MapCharacters[1].ID, 1)
	assert.Equal(t, mgr.MapCharacters[1].Name, "主角")
	assert.Equal(t, mgr.MapCharacters[1].HP, 100)
	assert.Equal(t, mgr.MapCharacters[1].Attack, 30)

	t.Logf("Test_LoadCharacterMgr OK")
}

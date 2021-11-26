package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoadCharacterMgr(t *testing.T) {
	mgr, err := LoadCharacterMgr("../gamedata/character.xlsx")
	assert.NoError(t, err)
	assert.Equal(t, len(mgr.MapCharacters), 41)

	assert.Equal(t, mgr.MapCharacters[1].ID, 1)
	assert.Equal(t, mgr.MapCharacters[1].Name, "绿头怪")
	assert.Equal(t, mgr.MapCharacters[1].HP, 50)
	assert.Equal(t, mgr.MapCharacters[1].Attack, 20)
	assert.Equal(t, mgr.MapCharacters[1].Defence, 1)
	assert.Equal(t, mgr.MapCharacters[1].Gold, 1)
	assert.Equal(t, mgr.MapCharacters[1].Exp, 1)

	t.Logf("Test_LoadCharacterMgr OK")
}

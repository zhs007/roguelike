package game

import "errors"

var (
	// ErrInvalidCharacterFiles - invalid character files
	ErrInvalidCharacterFiles = errors.New("invalid character files")
	// ErrInvalidCharacterID - invalid characterID
	ErrInvalidCharacterID = errors.New("invalid characterID")
)

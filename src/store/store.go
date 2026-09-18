package store

import (
	"switcher/game"
	"switcher/game/logic"
)

type GameStore interface {
	Get(key string) (*logic.Game, error)
	Set(key string, value *logic.Game) error
	Delete(key string) error
}

type UserStore interface {
	Get(key string) (*game.Player, error)
	Set(key string, value *game.Player) error
	Delete(key string) error
}

var ErrNotFound = &StoreError{"not found"}

type StoreError struct {
	Message string
}

func (e *StoreError) Error() string {
	return e.Message
}

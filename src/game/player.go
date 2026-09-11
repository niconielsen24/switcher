package game

import "uuid"

type Player struct {
	Name string
	Id   uuid.UUID
}

func NewPlayer(name string) *Player {
	return &Player{Name: name, Id: uuid.New()}
}

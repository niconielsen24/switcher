package game

import "uuid"

type Player struct {
	Name string    `json:"name"`
	Id   uuid.UUID `json:"id"`
}

func NewPlayer(name string) *Player {
	return &Player{Name: name, Id: uuid.New()}
}

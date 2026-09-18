package game

import (
	"math/rand"
	"uuid"
)

type FigureCard struct {
	Id    uuid.UUID `json:"id"`
	Shape Shape     `json:"shape"`
}

func NewFigureCard() *FigureCard {
	index := rand.Intn(len(ShapeList))
	s := ShapeList[index]
	return &FigureCard{Id: uuid.New(), Shape: s}
}

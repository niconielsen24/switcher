package game

import "uuid"

type FigureCardDeck = []FigureCard

type FigureCard struct {
	Id    uuid.UUID
	shape Shape
}

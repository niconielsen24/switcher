package game

import "uuid"

const (
	MOV_DECK_SIZE = 10
	FIG_DECK_SIZE = 3
)

type Player struct {
	Name         string                  `json:"name"`
	Id           uuid.UUID               `json:"id"`
	MovementDeck map[string]MovementCard `json:"movement_deck"`
	FigureDeck   map[string]FigureCard   `json:"figure_deck"`
}

func NewPlayer(name string) *Player {
	p := &Player{
		Name:         name,
		Id:           uuid.New(),
		MovementDeck: createMovementDeck(),
		FigureDeck:   createFigureDeck(),
	}

	return p
}

func createMovementDeck() map[string]MovementCard {
	cards := map[string]MovementCard{}
	for range MOV_DECK_SIZE {
		card := NewMovementCard()
		cards[card.Id.String()] = *card
	}
	return cards
}

func createFigureDeck() map[string]FigureCard {
	cards := map[string]FigureCard{}
	for range FIG_DECK_SIZE {
		card := NewFigureCard()
		cards[card.Id.String()] = *card
	}
	return cards
}

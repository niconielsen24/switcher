package store

import "switcher/game/logic"

type InmemGameStore struct {
	data map[string]*logic.Game
}

func NewInmemGameStore() *InmemGameStore {
	return &InmemGameStore{
		data: make(map[string]*logic.Game),
	}
}

func (s *InmemGameStore) Get(key string) (*logic.Game, error) {
	if game, ok := s.data[key]; ok {
		return game, nil
	}
	return nil, ErrNotFound
}

func (s *InmemGameStore) Set(key string, value *logic.Game) error {
	s.data[key] = value
	return nil
}

func (s *InmemGameStore) Delete(key string) error {
	delete(s.data, key)
	return nil
}

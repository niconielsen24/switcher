package store

import "switcher/game"

type InmemUserStore struct {
	data map[string]*game.Player
}

func NewInmemUserStore() *InmemUserStore {
	return &InmemUserStore{
		data: make(map[string]*game.Player),
	}
}

func (s *InmemUserStore) Get(key string) (*game.Player, error) {
	if player, ok := s.data[key]; ok {
		return player, nil
	}
	return nil, ErrNotFound
}

func (s *InmemUserStore) Set(key string, value *game.Player) error {
	s.data[key] = value
	return nil
}

func (s *InmemUserStore) Delete(key string) error {
	delete(s.data, key)
	return nil
}

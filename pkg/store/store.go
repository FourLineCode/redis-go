package store

import "errors"

type Store struct {
	db map[string]string
}

func New() *Store {
	return &Store{
		db: map[string]string{},
	}
}

func (s *Store) Set(key, value string) bool {
	s.db[key] = value

	return true
}

func (s *Store) Get(key string) (string, error) {
	val, ok := s.db[key]

	if !ok {
		return "", errors.New("key doesn't exist")
	}

	return val, nil
}

func (s *Store) Delete(key string) error {
	_, ok := s.db[key]

	if !ok {
		return errors.New("key doesn't exist")
	}

	delete(s.db, key)
	return nil
}

package store

import "errors"

func (s *Store) HSet(key, field, value string) bool {
	_, ok := s.hashDb[key]
	if !ok {
		s.hashDb[key] = Map{}
	}

	s.hashDb[key][field] = value

	return true
}

func (s *Store) HGet(key, field string) (string, error) {
	k, ok := s.hashDb[key]
	if !ok {
		return "", errors.New("key doesn't exist")
	}

	val, found := k[field]
	if !found {
		return "", errors.New("field doesn't exist on given key")
	}

	return val, nil
}

func (s *Store) HGetAll(key string) (Map, error) {
	val, ok := s.hashDb[key]
	if !ok {
		return Map{}, errors.New("key doesn't exist")
	}

	return val, nil
}

func (s *Store) HDelete(key string) error {
	_, ok := s.hashDb[key]

	if !ok {
		return errors.New("key doesn't exist")
	}

	delete(s.hashDb, key)
	return nil
}

package store

import "errors"

func (s *Store) LPush(key, value string) bool {
	_, ok := s.listDb[key]

	if !ok {
		s.listDb[key] = List{}
	}

	s.listDb[key] = append([]string{value}, s.listDb[key]...)

	return true
}

func (s *Store) RPush(key, value string) bool {
	_, ok := s.listDb[key]

	if !ok {
		s.listDb[key] = List{}
	}

	s.listDb[key] = append(s.listDb[key], value)

	return true
}

func (s *Store) LPop(key string) (string, error) {
	list, ok := s.listDb[key]

	if !ok {
		return "", errors.New("key doesn't exist")
	}

	if len(list) <= 0 {
		return "(nil)", nil
	}

	var value string
	value, s.listDb[key] = list[0], list[1:]

	return value, nil
}

func (s *Store) RPop(key string) (string, error) {
	list, ok := s.listDb[key]

	if !ok {
		return "", errors.New("key doesn't exist")
	}

	if len(list) <= 0 {
		return "(nil)", nil
	}

	var value string
	value, s.listDb[key] = list[len(list)-1], list[:len(list)-1]

	return value, nil
}

func (s *Store) LIndex(key string, index int) (string, error) {
	list, ok := s.listDb[key]

	if !ok {
		return "", errors.New("key doesn't exist")
	}

	if index >= len(list) || index < 0 {
		return "", errors.New("index out of range")
	}

	if len(list) <= 0 {
		return "(nil)", nil
	}

	value := list[index]

	return value, nil
}

func (s *Store) LRange(key string, start, end int) ([]string, error) {
	list, ok := s.listDb[key]

	if !ok {
		return []string{}, errors.New("key doesn't exist")
	}

	if start >= len(list) || start < 0 {
		return []string{}, errors.New("start index out of range")
	}

	if end < start || end > len(list) {
		return []string{}, errors.New("end index out of range")
	}

	if len(list) <= 0 {
		return []string{}, nil
	}

	value := list[start:end]

	return value, nil
}

func (s *Store) LDelete(key string) error {
	_, ok := s.listDb[key]

	if !ok {
		return errors.New("key doesn't exist")
	}

	delete(s.listDb, key)
	return nil
}

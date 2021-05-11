package store

type Map map[string]string
type List []string

type Store struct {
	db     map[string]string
	hashDb map[string]Map
	listDb map[string]List
}

func New() *Store {
	return &Store{
		db:     map[string]string{},
		hashDb: map[string]Map{},
		listDb: map[string]List{},
	}
}

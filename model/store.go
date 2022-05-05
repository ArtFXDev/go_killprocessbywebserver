package model

// type Store interface{}

type Store struct {
	NimbyStatus *NimbyStatus
}

func NewStoreInstance() *Store {
	return &Store{NewNimbyStatus()}
}

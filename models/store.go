package models

import (
	"sync"
)

// type Store interface{}

type Store struct {
	NimbyStatus *NimbyStatus
}

func newStoreInstance() *Store {
	return &Store{NewNimbyStatus()}
}

var lock = &sync.Mutex{}

var storeInstance *Store

func GetStoreInstance() *Store {
	if storeInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if storeInstance == nil {
			storeInstance = newStoreInstance()
		}
	}
	return storeInstance
}

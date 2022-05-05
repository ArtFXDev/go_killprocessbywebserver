package store

import (
	"log"
	"sync"

	"github.com/OlivierArgentieri/go_killprocess/model"
)

var lock = &sync.Mutex{}

// type Store interface{}

type Store struct {
	NimbyStatus *model.NimbyStatus
}

var storeInstance *Store

func newStoreInstance() *Store {
	return &Store{model.NewNimbyStatus()}
}

func GetInstance() *Store {
	if storeInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if storeInstance == nil {
			log.Printf("Create instance")
			storeInstance = newStoreInstance()
		} else {
			log.Printf("already created")
		}
	} else {
		log.Printf("already created")
	}

	return storeInstance
}

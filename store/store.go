package store

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Store struct {
}

var storeInstance *Store

func getInstance() *Store {
	if storeInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if storeInstance == nil {
			fmt.Println("Create instance")
			storeInstance = &Store{}
		} else {
			fmt.Println("already created")
		}
	} else {
		fmt.Println("already created")
	}

	return storeInstance
}

package store

import (
	"log"
	"sync"

	"github.com/OlivierArgentieri/go_killprocess/model"
)

var lock = &sync.Mutex{}

var storeInstance *model.Store

func GetInstance() *model.Store {
	if storeInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if storeInstance == nil {
			log.Printf("Create instance")
			storeInstance = model.NewStoreInstance()
		} else {
			log.Printf("already created")
		}
	} else {
		log.Printf("already created")
	}

	return storeInstance
}

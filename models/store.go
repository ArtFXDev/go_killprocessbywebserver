package models

import (
	"sync"
)

// type Store interface{}

type Store struct {
	NimbyStatus   *NimbyStatus
	loopIsRunning bool
}

func newStoreInstance() *Store {
	newStatus := NewNimbyStatus()
	return &Store{newStatus, false}
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

func UpdateInstance(ns *NimbyStatus) {
	if ns.Value != nil {
		GetStoreInstance().NimbyStatus.SetValue(*ns.Value)
	}

	if ns.Mode != nil {
		GetStoreInstance().NimbyStatus.SetMode(*ns.Mode)
	}

	if ns.Reason != nil {
		GetStoreInstance().NimbyStatus.SetReason(*ns.Reason)
	}
}

package store

import (
	"errors"
	"sync"
)

type SafeMap struct {
	rwmutex sync.RWMutex
	kvstore map[string]string
}


func (sm *SafeMap) Create (key, value string) error{
	sm.rwmutex.Lock()
	defer sm.rwmutex.Unlock()
	_, alreadyExists := sm.kvstore[key]
	if alreadyExists {
		return errors.New("key already exists")
	}
	sm.kvstore[key] = value
	return nil
}

func (sm *SafeMap) Retrieve(key string) (error, string) {
	sm.rwmutex.RLock()
	defer sm.rwmutex.RUnlock()
	if val, keyExists := sm.kvstore[key]; keyExists {
		return nil, val
	} else {
		return errors.New("key does not exist"), ""
	}
}

func (sm *SafeMap) Update(key, value string) error {
	sm.rwmutex.Lock()
	defer sm.rwmutex.Unlock()
	if _, keyExists := sm.kvstore[key]; keyExists {
		sm.kvstore[key] = value
		return nil
	}  else {
		return errors.New("key does not exist")
	}
}

func (sm *SafeMap) Delete(key string) error {
	sm.rwmutex.RLock()
	defer sm.rwmutex.RUnlock()
	if _, keyExists := sm.kvstore[key]; keyExists {
		delete(sm.kvstore, key)
		return nil
	}  else {
		return errors.New("key does not exist")
	}
}

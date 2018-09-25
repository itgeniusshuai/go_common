package common

import "sync"

type Map struct{
	data map[interface{}]interface{}

	lock sync.RWMutex
}

func (this * Map)Put(k interface{},v interface{}){
	defer this.lock.Unlock()
	this.lock.Lock()
	this.data[k]=v
}

func (this * Map)Get(k interface{}) interface{}{
	defer this.lock.RUnlock()
	this.lock.RLock()
	return this.data[k]
}

func (this * Map)Length()int{
	return len(this.data)
}
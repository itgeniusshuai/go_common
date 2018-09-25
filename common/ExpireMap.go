package common

import (
	"sync"
	"time"
)

type ExpireMap struct{
	data map[interface{}]interface{}

	expire map[interface{}]*time.Time

	lock sync.RWMutex
}

func (expireMap *ExpireMap)Put(key interface{},v interface{}){
	defer expireMap.lock.Unlock()
	expireMap.lock.Lock()
	expireMap.makeSureDataAndExpire()
	expireMap.data[key] = v
}

func (expireMap *ExpireMap)PutWithExpire(key interface{},v interface{},duration time.Duration){
	defer expireMap.lock.Unlock()
	expireMap.lock.Lock()
	expireMap.makeSureDataAndExpire()
	expireMap.data[key] = v
	t := time.Now().Add(duration)
	expireMap.expire[key] = &t
}

// Get的时候获取时间
func (expireMap *ExpireMap)Get(key interface{}) interface{}{
	defer expireMap.lock.Unlock()
	expireMap.lock.Lock()
	expireTime := expireMap.expire[key]
	if expireTime == nil || expireTime.After(time.Now()){
		return expireMap.data[key]
	}
	expireMap.Delete(key)
	return nil
}

// 删除并删除对应的过期时间
func (expireMap *ExpireMap)Delete(key interface{}){
	defer expireMap.lock.Unlock()
	expireMap.lock.Lock()
	delete(expireMap.data,key)
	delete(expireMap.expire,key)
}

func (expireMap *ExpireMap)makeSureDataAndExpire(){
	if expireMap.data == nil{
		expireMap.data = make(map[interface{}]interface{})
	}
	if expireMap.expire == nil{
		expireMap.expire = make(map[interface{}]*time.Time)
	}
}


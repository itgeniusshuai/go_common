package common

import (
	"hash/crc32"
	"sync"
)
/**
	目前不支持扩容
 */

const DEFAULT_LINK_SIZE = 16

type concurrentHashMap struct {
	// 桶加链表
	links []*MapLink

}

type MapLink struct {
	lock *sync.RWMutex
	root *Node
	tail *Node
}

type Node struct{
	hash int
	key string
	value interface{}
	next *Node
	pre *Node

}


func InitConcurrentHashMap() *concurrentHashMap{
	links := make([]*MapLink,DEFAULT_LINK_SIZE)
	for i := 0; i < DEFAULT_LINK_SIZE; i++{
		lock := new(sync.RWMutex)
		link := &MapLink{lock:lock}
		links[i] = link
	}
	return &concurrentHashMap{links:links}
}

// 插入值
func (this *concurrentHashMap)Put(key string, value interface{}){
	// 计算hash值
	link := this.links[hash(key,DEFAULT_LINK_SIZE)]
	link.Put(key,value)
}

// 取出值
func (this *concurrentHashMap)Get(key string) interface{}{
	link := this.links[hash(key,DEFAULT_LINK_SIZE)]
	return link.Get(key)
}

// 移除值
func (this *concurrentHashMap)Remove(key string) interface{}{
	return this.links[hash(key,DEFAULT_LINK_SIZE)].Remove(key)
}

// 是否包含
func (this *concurrentHashMap)Contain(key string) bool{
	return this.links[hash(key,DEFAULT_LINK_SIZE)].Contain(key)
}

// 所有key
func (this *concurrentHashMap)KeySet() []string{
	keys := make([]string,0)
	for _,e := range this.links{
		var node *Node
		if e == nil{
			continue
		}
		node = e.root
		for node != nil{
			keys = append(keys, node.key)
			node = node.next
		}
	}
	return keys;
}

// 所有value
func (this *concurrentHashMap)Values() []interface{}{
	keys := make([]interface{},0)
	for _,e := range this.links{
		var node *Node
		if e == nil{
			continue
		}
		node = e.root
		for node != nil{
			keys = append(keys, node.value)
			node = node.next
		}
	}
	return keys;
}

// 移除元素
func (this *MapLink)Remove(key string) interface{}{
	defer this.lock.Unlock()
	this.lock.Lock()
	node := this.root
	for(node != nil){
		if node.key == key{
			// 头节点
			if node == this.root{
				this.root = node.next
				node.next = nil
			}else if node == this.tail{
				this.tail = node.pre
				node.pre.next = nil
			}else {
				node.pre.next = node.next
				node.next.pre =node.pre
				node.pre = nil
				node.next = nil
			}
			return node.value
		}
		node = node.next
	}
	return nil
}

func (this *MapLink)Contain(key string) bool{
	node := this.root
	for(node != nil){
		if node.key == key{
			return true
		}
		node = node.next
	}
	return false
}


func (this *MapLink)Get(key string) interface{}{
	defer this.lock.Unlock()
	this.lock.Lock()
	node := this.root
	for(node != nil){
		if node.key == key{
			return node.value
		}
		node = node.next
	}
	return nil
}

func (this *MapLink)Put(key string, v interface{}){
	defer this.lock.Unlock()
	this.lock.Lock()
	node := this.root
	for(node != nil){
		if node.key == key{
			node.value = v
		}
		node = node.next
	}
	newNode := Node{key:key,value:v}
	lastTail := this.tail
	if lastTail == nil{
		this.root = &newNode
		this.tail = &newNode
	}else{
		this.tail.next = &newNode
		newNode.pre = this.tail
	}

}

func hash(key string, length int) int{
	hashV := crc32.ChecksumIEEE([]byte(key))
	hash := int(hashV) & (length >> 1)
	return hash
}


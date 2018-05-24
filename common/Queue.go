package common

import "sync"

type Queue struct{
	head *node
	tail *node

	lock sync.RWMutex
}

type node struct{
	v interface{}

	next *node
}

func (q *Queue)Push(v interface{}){
	defer q.lock.Unlock()
	q.lock.Lock()
	node := node{v:v}
	// 首次入队，头尾都指向头
	if q.head == nil{
		q.head = &node
		q.tail = &node
		return
	}
	// 非首次
	q.tail.next =&node
	q.tail = &node
}

func (q *Queue)Pop()interface{}{
	defer q.lock.Unlock()
	q.lock.Lock()
	if q.head == nil{
		return nil
	}
	oldHead := q.head
	q.head = oldHead.next
	oldHead.next = nil
	return oldHead.v
}
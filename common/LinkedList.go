package common

import (
	"fmt"
	"errors"
	"sync"
)

type LinkedList struct {
	head *lNode
	tail *lNode
	length int64

	lock sync.RWMutex
}

type lNode struct {
	Value interface{}
	next *lNode
	pre *lNode
}

func CreateLinkedList() LinkedList{
	return LinkedList{length:0}
}

func (this *LinkedList)IsEmpty() bool{
	return this.length == 0
}
func (this *LinkedList)Lpush(v interface{})(error){
	defer this.lock.Unlock()
	this.lock.Lock()
	if v == nil{
		return errors.New("push an empty element")
	}
	node := &lNode{Value:v}
	if(this.IsEmpty()){
		this.tail = node
	}else {
		this.head.pre = node
		node.next=this.head
	}
	this.head = node
	this.length += 1
	return nil
}
func (this *LinkedList)Lpop()(interface{}){
	defer this.lock.Unlock()
	this.lock.Lock()
	if(this.IsEmpty()){
		return nil
	}
	this.length-=1
	v := this.head.Value

	if (this.head == this.tail){
		this.head = nil
		this.tail = nil
		return v
	}
	this.head = this.head.next
	this.head.pre.next = nil
	this.head.pre = nil
	return v
}

func (this *LinkedList)Rpush(v interface{})(error){
	defer this.lock.Unlock()
	this.lock.Lock()
	node := &lNode{Value:v}
	if(this.IsEmpty()){
		this.head = node
	}else {
		this.tail.next = node
		node.pre = this.tail
	}
	this.tail = node
	this.length += 1
	return nil
}

func (this *LinkedList)Rpop()(interface{}){
	if(this.IsEmpty()){
		return nil
	}
	this.length-=1
	v := this.tail.Value

	if (this.head == this.tail){
		this.head = nil
		this.tail = nil
		return v
	}
	this.tail = this.tail.pre
	this.tail.next.pre = nil
	this.tail.next = nil
	return v
}

func (this *LinkedList)PrintList(){
	if this.IsEmpty(){
		fmt.Println("<nil>")
		return
	}
	node := this.head
	fmt.Print("[")
	for node != this.tail{
		fmt.Print(node.Value)
		fmt.Print(",")
		node = node.next
	}
	fmt.Print(this.tail.Value)
	fmt.Print("]")
}

func (this *LinkedList)LpopN(n int64)([]interface{}){
	defer this.lock.Unlock()
	this.lock.Lock()
	if(this.length == 0 || n <= 0){
		return nil
	}
	if(this.length < n){
		n = this.length
	}
	this.length = this.length-n
	res := make([]interface{},n)
	var i int64
	node := this.head
	res[0] = node.Value
	for i = 1; i < n; i++{
		res[i] = node.next.Value
		node = node.next
		node.pre.next = nil
		node.pre = nil
	}
	this.head = node
	if(node == nil){
		this.tail = nil
	}
	return res
}

func (this *LinkedList)RpopN(n int64)([]interface{}){
	defer this.lock.Unlock()
	this.lock.Lock()
	if(this.length == 0 || n <= 0){
		return nil
	}
	if(this.length < n){
		n = this.length
	}
	this.length = this.length-n
	res := make([]interface{},n)
	var i int64
	node := this.tail
	res[0] = node.Value
	for i = 1; i < n; i++{
		res[i] = node.pre.Value
		node = node.pre
		node.next.pre = nil
		node.next = nil
	}
	this.tail = node
	if(node == nil){
		this.head = nil
	}
	return res
}

func (this *LinkedList)Add(v interface{}){
	this.Rpush(v)
}

func (this *LinkedList)Set(index int64, v interface{}) error{
	defer this.lock.Unlock()
	this.lock.Lock()
	if this.length <= index{
		return errors.New("index of slide")
	}
	var i int64 = 0
	var nodeP *lNode = this.head
	for ;i < index; i++{
		nodeP = nodeP.next
	}
	nodeP.Value = v
	return nil
}

func (this *LinkedList)Remove(index int64) (interface{},error){
	defer this.lock.Unlock()
	this.lock.Lock()
	if this.length <= index{
		return nil,errors.New("index of slide")
	}
	var i int64 = 0
	var nodeP *lNode = this.head
	for ;i < index-1; i++{
		nodeP = nodeP.next
	}
	rNode := nodeP.next
	nodeP.next = rNode.next
	nodeP.next.pre = nodeP
	rNode.pre =nil
	rNode.next = nil
	return rNode.Value,nil
}


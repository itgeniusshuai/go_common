package common

import (
	"sync"
)

type Stack struct{
	data []interface{}

	lock sync.RWMutex
}

func GetStack() Stack{
	return Stack{}
}

func (this *Stack)IsEmpty()bool{
	return len(this.data) == 0;
}

func (this *Stack)Push(v interface{}){
	defer this.lock.Unlock()
	this.lock.Lock()
	this.data = append(this.data, v)
}

func (this *Stack)Pop()interface{}{
	defer this.lock.Unlock()
	this.lock.Lock()
	var length = len(this.data)
	if length == 0{
		return nil
	}
	var v = this.data[length-1]
	endIndex := Max(length-2,0)
	this.data = this.data[:endIndex]
	return v
}

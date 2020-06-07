package stack

import "sync"

type ArrayStack struct {
	Arr  []interface{}
	lock sync.Mutex
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{lock: sync.Mutex{}, Arr: make([]interface{}, 0)}
}

func (as *ArrayStack) Peek() interface{} {
	n := len(as.Arr)
	if n == 0 {
		return -1
	}
	return as.Arr[n-1]
}

func (as *ArrayStack) Push(value interface{}) {
	as.lock.Lock()
	defer as.lock.Unlock()
	as.Arr = append(as.Arr, value)
}

func (as *ArrayStack) Pop() interface{} {
	as.lock.Lock()
	defer as.lock.Unlock()
	n := len(as.Arr)
	if n == 0 {
		return -1
	}
	ret := as.Arr[n-1]
	as.Arr = as.Arr[:n-1]
	return ret
}

package stack

type Stack interface {
	Peek() interface{}
	Push(value interface{})
	Pop() interface{}
}

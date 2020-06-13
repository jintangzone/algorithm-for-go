package stack

//type ArrayStack struct {
//	arr []interface{}
//}
//
//func NewArrayStack() *ArrayStack {
//	return &ArrayStack{arr: make([]interface{}, 0)}
//}
//
//func (as *ArrayStack) Push(v interface{}) {
//	as.arr = append(as.arr, v)
//}
//
//func (as *ArrayStack) Peek() interface{} {
//	return as.arr[len(as.arr)-1]
//}
//
//func (as *ArrayStack) Pop() interface{} {
//	node := as.arr[len(as.arr)-1]
//	as.arr = as.arr[:len(as.arr)-1]
//	return node
//}
//
//func (as *ArrayStack) Size() int {
//	return len(as.arr)
//}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	//stk := NewArrayStack()
	stk := make([]rune, 0)
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			//stk.Push(c)
			stk = append(stk, c)
		} else {
			n := len(stk)
			if n == 0 {
				return false
			}
			//top := stk.Peek().(rune)
			top := stk[n-1]
			if top == '{' && c != '}' ||
				top == '[' && c != ']' ||
				top == '(' && c != ')' {
				return false
			}
			//stk.Pop()
			stk = stk[:n-1]
		}
	}
	return len(stk) == 0
}

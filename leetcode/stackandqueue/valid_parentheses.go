package stackandqueue

import "algorithm-for-go/base/structure/stack"

func IsValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stk := new(stack.ArrayStack)
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stk.Push(c)
		} else {
			if len(stk.Arr) == 0 {
				return false
			}
			top := stk.Peek().(rune)
			if top == '{' && c != '}' ||
				top == '[' && c != ']' ||
				top == '(' && c != ')' {
				return false
			}
			stk.Pop()
		}
	}
	return len(stk.Arr) == 0
}

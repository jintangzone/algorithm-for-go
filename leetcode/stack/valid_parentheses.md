# 20.有效的括号

### 原题

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

- 左括号必须用相同类型的右括号闭合。
- 左括号必须以正确的顺序闭合。

注意空字符串可被认为是有效字符串。

来源：力扣（LeetCode）

[链接：https://leetcode-cn.com/problems/valid-parentheses](https://leetcode-cn.com/problems/valid-parentheses)

### 使用栈

```go
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
```

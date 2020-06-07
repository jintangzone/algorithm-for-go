## 141.环形链表

### 原题

给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

来源：力扣（LeetCode）

[链接：https://leetcode-cn.com/problems/linked-list-cycle](https://leetcode-cn.com/problems/linked-list-cycle)

### 快慢指针法

```go
func HasCycle(link *linkedlist.LinkedList) bool {
	fast := link.Head
	slow := fast
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
```

思路分析：

```
使用两个指针，一个slow，一个fast
 
let：
slow = slow.Next
fast = fast.Next.Next
 
slow走一步，fast走两步，如果他们能够相遇，则说明有环
 
如果两个都等于null时，则没有环
```

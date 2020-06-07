# 24.两两交换链表中的节点

### 原题

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:
```
给定 1->2->3->4, 你应该返回 2->1->4->3.
```

来源：力扣（LeetCode）

[链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs](https://leetcode-cn.com/problems/swap-nodes-in-pairs)

### 借空节点的三指针法

```go
func SwapNode(link *linkedlist.LinkedList) {
	if link.Head == nil || link.Head.Next == nil {
		return
	}
	pre := new(linkedlist.LinkNode)
	pre.Next = link.Head
	link.Head = link.Head.Next
	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := a.Next
		pre.Next, b.Next, a.Next = b, a, b.Next
		pre = a
	}
}
```
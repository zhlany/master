package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 写一个反转链表的函数
// 反转链表
func reverseList(head *ListNode) *ListNode {
	// 前一个节点
	var prev *ListNode
	// 当前节点
	curr := head
	// 当当前节点不为空时
	for curr != nil {
		// 下一个节点
		next := curr.Next
		// 当前节点的下一个节点指向前一个节点
		curr.Next = prev
		// 前一个节点指向当前节点
		prev = curr
		// 当前节点指向下一个节点
		curr = next
	}
	// 返回反转后的头节点
	return prev
}

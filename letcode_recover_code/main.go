package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// main 函数演示了如何反转一个链表并打印结果。
// 创建了一个包含5个节点的链表，然后调用 reverse 函数进行反转，
// 最后遍历并打印反转后的链表节点值。
func main() {
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next = &ListNode{Value: 4}
	head.Next.Next.Next.Next = &ListNode{Value: 5}

	newHead := reverse(head)
	for newHead != nil {
		fmt.Print(newHead.Value, " ")
		newHead = newHead.Next
	}
}

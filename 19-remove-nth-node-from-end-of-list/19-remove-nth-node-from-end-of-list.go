package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}
	if n == 0 {
		return head
	}
	var length int
	var current *ListNode = head
	for current != nil {
		length++
		current = current.Next
	}
	if length == n {
		return head.Next
	}
	current = head
	for i := 0; i < length-n-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	return head
}

func printList(head *ListNode) {
	var current *ListNode = head
	for current != nil {
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// 1->2->3->4->5, and n = 2.
	// After removing the second node from the end, the linked list becomes 1->2->3->5.
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	printList(head)
	printList(removeNthFromEnd(head, 2))
}

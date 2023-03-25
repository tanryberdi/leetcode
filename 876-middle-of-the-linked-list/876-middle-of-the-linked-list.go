package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var length int
	for node := head; node != nil; node = node.Next {
		length++
	}
	for i := 0; i < length/2; i++ {
		head = head.Next
	}
	return head
}

func main() {
	// 1->2->3->4->5
	// 1->2->3->4->5->6
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	fmt.Println(middleNode(head).Val)
}

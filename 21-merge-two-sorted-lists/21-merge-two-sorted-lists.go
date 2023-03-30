package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode

	for list1 != nil && list2 != nil {
		var min *ListNode
		if list1.Val < list2.Val {
			min = list1
			list1 = list1.Next
		} else {
			min = list2
			list2 = list2.Next
		}
		if head == nil {
			head = min
		} else {
			tail.Next = min
		}
		tail = min
	}

	if list1 != nil {
		if head == nil {
			head = list1
		} else {
			tail.Next = list1
		}
	}

	if list2 != nil {
		if head == nil {
			head = list2
		} else {
			tail.Next = list2
		}
	}

	return head
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	fmt.Println(mergeTwoLists(list1, list2))
}

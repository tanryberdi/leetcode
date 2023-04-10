package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		var nextQueue []*Node
		for i := 0; i < len(queue); i++ {
			if i < len(queue)-1 {
				queue[i].Next = queue[i+1]
			} else {
				queue[i].Next = nil
			}
			if queue[i].Left != nil {
				nextQueue = append(nextQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				nextQueue = append(nextQueue, queue[i].Right)
			}
		}
		queue = nextQueue
	}
	return root
}

func main() {
	root := &Node{}
	fmt.Println(connect(root))
}

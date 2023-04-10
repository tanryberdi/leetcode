package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	subRoot := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isSubtree(root, subRoot))
}

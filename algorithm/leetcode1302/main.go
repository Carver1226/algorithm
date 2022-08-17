package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	max := deepestLevel(root)
	ans := deepestLevelSum(root, max, 1)
	return ans
}

func deepestLevelSum(root *TreeNode, max int, depth int) int {
	if root != nil {
		if depth == max {
			return root.Val
		}
		return deepestLevelSum(root.Left, max, depth + 1) + deepestLevelSum(root.Right, max, depth + 1)
	}
	return 0
}

func deepestLevel(root *TreeNode) int {
	if root != nil {
		return max(deepestLevel(root.Left), deepestLevel(root.Right)) + 1
	}
	return 0
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
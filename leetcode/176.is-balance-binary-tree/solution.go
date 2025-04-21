package solution

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func IsBalance(root *Node) bool {
	if root == nil {
		return true
	}
	return abs(maxDepth(root.Left)-maxDepth(root.Right)) <= 1 && IsBalance(root.Left) && IsBalance(root.Right)
}

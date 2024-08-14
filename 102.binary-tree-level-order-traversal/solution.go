package solution

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	currentLevel := []*TreeNode{root}
	for len(currentLevel) > 0 {
		nextLevel := []*TreeNode{}
		var temp []int
		for _, node := range currentLevel {
			temp = append(temp, node.Value)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		currentLevel = nextLevel
		result = append(result, temp)
	}
	return result
}

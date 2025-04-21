package solution

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// 经典的层序遍历
func LevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	// 初始化的时候就把root赋值给当前层
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
	fmt.Println(result)
	return result
}

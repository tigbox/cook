package solution

type Node struct {
	Value int
	Next  *Node
}

func ReverseNodeList(root *Node) *Node {
	if root == nil {
		return nil
	}
	var previous *Node
	current := root
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}

// 第一步：声明previous，返回previous
func ReverseNodeList1(root *Node) *Node {
	if root == nil {
		return nil
	}
	var previous *Node
	return previous
}

// 第二步：初始化current，遍历中current游标下去
func ReverseNodeList2(root *Node) *Node {
	if root == nil {
		return nil
	}
	var previous *Node
	current := root
	for current != nil {
		next := current.Next

		current = next
	}

	return previous
}

// 第三步：currentNext指向previous，current赋值给previous
func ReverseNodeList3(root *Node) *Node {
	if root == nil {
		return nil
	}
	var previous *Node
	current := root
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}

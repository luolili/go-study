package main

func main() {

}

//迭代
func maxDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	var level = 0
	for len(q) > 0 {
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}

		}
		level++
	}
	return level
}

// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

//go 没有三木运算符
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

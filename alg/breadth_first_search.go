package main

import "container/list"

func main() {

}

func breadthFirstSearch(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := list.New()
	q.PushFront(root)
	for q.Len() > 0 {
		var curLevel []int
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			node := q.Remove(q.Back()).(*TreeNode)
			curLevel = append(curLevel, node.Value)
			if node.Left != nil {
				q.PushFront(node.Left)
			}
			if node.Right != nil {
				q.PushFront(node.Right)
			}
		}
		res = append(res, curLevel)
	}
	return res
}

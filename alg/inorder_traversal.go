package main

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func inorderTraversal(root *TreeNode) []int {

	var res = make([]int, 0)
	return buildTree(root, &res)
}

func buildTree(root *TreeNode, res *[]int) []int {

	if root != nil {
		if root.Left != nil {
			buildTree(root.Left, res)
		}
		*res = append(*res, root.Value)
		if root.Right != nil {
			buildTree(root.Right, res)
		}
	}
	return *res
}

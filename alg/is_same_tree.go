package main

func main() {

}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Value != q.Value {
		return false
	}
	return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}

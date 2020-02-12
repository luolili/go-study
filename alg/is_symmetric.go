package main

func main() {

}
func isMirror(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Value == q.Value && isMirror(p.Right, q.Left) && isMirror(p.Left, q.Right)
}

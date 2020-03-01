package main

func main() {

}

var balanced = true

func isBalanced(root *TreeNode) bool {
	height(root)
	return balanced
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
		balanced = false
	}
	return max(leftHeight, rightHeight) + 1
}

func helper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	var leftHeight, leftFlag = helper(root.Left)
	var rightHeight, rightFlag = helper(root.Right)
	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
		return 0, false
	}
	var level int
	if leftHeight > rightHeight {
		level = leftHeight
	} else {
		level = rightHeight
	}
	return level + 1, leftFlag && rightFlag

}

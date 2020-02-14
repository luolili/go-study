package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func deleteNode(head *ListNode, val int) *ListNode {

	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	curNode := head
	for curNode != nil && curNode.Next != nil {
		if curNode.Next.Val == val {
			curNode.Next = curNode.Next.Next
		}
		curNode = curNode.Next
	}
	return head
}

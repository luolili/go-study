package main

func main() {

}
func reversePrint(head *ListNode) []int {
	len := 0
	var preNode *ListNode = nil
	curNode := head
	for curNode != nil {
		tmp := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = tmp
		len++
	}

	res := make([]int, len)
	i := 0
	for preNode != nil {
		res[i] = preNode.Val
		preNode = preNode.Next
		i++
	}
	return res
}
func reverseLinkedListV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := head.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

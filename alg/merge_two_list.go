package main

//输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
func main() {

}
func mergeTwoListV2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoListV2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListV2(l2.Next, l1)
		return l2
	}
}
func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	//注意这里的赋值 不能是 var dummyNode *ListNode & 取 var 地址
	var dummyNode = &ListNode{0, nil}
	var newList = dummyNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			dummyNode.Next = l1
			l1 = l1.Next
		} else {
			dummyNode.Next = l2
			l2 = l2.Next
		}
		dummyNode = dummyNode.Next
	}
	if l1 != nil {
		dummyNode.Next = l1
	}
	if l2 != nil {
		dummyNode.Next = l2
	}
	return newList.Next
}

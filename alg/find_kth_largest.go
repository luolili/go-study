package main

import "container/heap"

// 第k大的元素
type MinHeap []int

// 括号里面 制定 变量名称
func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 括号里面
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := old.Len() - 1
	x := old[n]
	*h = old[0:n]
	return x

}
func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return 0
	}
	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return heap.Pop(h).(int)
}
func main() {

}

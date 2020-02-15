package main

type OQueue struct {
	in  []int
	out []int
}

func Constructor() OQueue {

	return OQueue{
		make([]int, 0, 5),
		make([]int, 0, 5),
	}
}
func (this *OQueue) AppendTail(value int) {
	if value < 1 || value > 1000 {
		return
	} else {
		this.in = append(this.in[:], value)
	}
}
func (this *OQueue) DeleteHead() int {

	if len(this.in) == 0 {
		return -1
	}
	if len(this.in) == len(this.out) {
		return -1
	}
	out := this.in[len(this.out)]
	this.out = append(this.out[:], this.in[len(this.out)])
	return out
}
func main() {

}

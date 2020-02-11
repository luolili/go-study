package main

func main() {

}
func myPow(x float64, n int32) float64 {
	if x == 0 {
		return 0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var res = 1.0
	var cur = x
	for i := n; i > 0; i /= 2 {
		if i%2 == 1 {
			res = res * cur
		}
		cur = cur * cur
	}
	return res
}

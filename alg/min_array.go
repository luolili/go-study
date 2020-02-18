package main

func main() {

}
func minArray(numbers []int) int {
	len := len(numbers)
	for i := 0; i < len-1; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}

	}
	return numbers[0]
}

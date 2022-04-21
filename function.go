package main

import "fmt"

func sum(a int, b int) (int, bool) {
	if a < 0 || b < 0 {
		return 0, false
	}
	return a + b, true
}

func avg(nums ...int) float64 {
	total := 0
	for _, num := range nums {
		total += num
	}
	average := float64(total / len(nums))
	return average
}

func main() {
	total, ok := sum(0, 0)
	if ok {
		fmt.Println(total)
	} else {
		fmt.Println("Negative values are not allowed")
	}
	nums := []int{2, 3, 45, 6, 757, 67, 4545}
	average := avg(nums...)
	fmt.Println(average)
}

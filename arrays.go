package main

import "fmt"

func main() {
	var intArr [10]int
	fmt.Println(intArr)
	for i := 0; i < 10; i++ {
		intArr[i] = i + 1
	}
	fmt.Println(intArr)
	fmt.Println(len(intArr))
	fmt.Println(cap(intArr))
}

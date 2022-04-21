package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{10, 20, 30, 40, 50, 1, 23, 3, 2}
	slice2 := intSlice[1:3]
	slice3 := intSlice[0:4]
	fmt.Println(slice2)
	fmt.Println(slice3)

	sort.Ints(intSlice)
	fmt.Println(intSlice)

	slice4 := make([]int, 10)
	fmt.Println(slice4)

	slice5 := make([]float64, 5)
	fmt.Println(slice5)

	slice6 := make([]string, 5)
	fmt.Println(slice6)

	slice7 := make([]bool, 4)
	fmt.Println(slice7)
	/*
		fmt.Println(intSlice)
		fmt.Println(len(intSlice))
		fmt.Println(cap(intSlice))
		intSlice = append(intSlice, 20)
		fmt.Println(intSlice)
		fmt.Println(len(intSlice))
		fmt.Println(cap(intSlice))
		intSlice = append(intSlice, 70)
		fmt.Println(intSlice)
		fmt.Println(len(intSlice))
		fmt.Println(cap(intSlice))
		intSlice = append(intSlice, 890)
		fmt.Println(intSlice)
		fmt.Println(len(intSlice))
		fmt.Println(cap(intSlice))
		intSlice = append(intSlice, 790)
		fmt.Println(intSlice)
		fmt.Println(len(intSlice))
		fmt.Println(cap(intSlice))
		intSlice = append(intSlice, 490)
		fmt.Println(intSlice)
		fmt.Println(len(intSlice))
		fmt.Println(cap(intSlice))
		intSlice = append(intSlice, 4590)
		fmt.Println(intSlice)
		fmt.Println(len(intSlice))
		fmt.Println(cap(intSlice))*/

}

package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}
	j := 0
	for j < 20 {
		fmt.Println("J is ", j)
		j++
	}

	for {
		fmt.Println("This is in infinit loop", j)
		j++
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}
}

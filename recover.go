package main

import "fmt"

func makePanic() {
	panic("Panic Occured")
}

func main() {
	defer func() {
		r := recover()
		fmt.Println(r)
		if r != nil {
			fmt.Println("Recovered")
		}
	}()
	makePanic()
	fmt.Println("This is not executed")
}

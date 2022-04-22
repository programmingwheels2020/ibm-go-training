package main

import (
	"fmt"
	"time"
)

func myFunction(from string) {
	for i := 0; i < 1000; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go myFunction("Red")
	go myFunction("Blue")

	go func(msg string) {
		fmt.Println(msg)
	}("Green")

	time.Sleep(time.Second)
	fmt.Println("done")
}

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func myFunction(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go myFunction("Red")
	go myFunction("Blue")
	wg.Wait()

}

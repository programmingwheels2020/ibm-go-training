package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var wg sync.WaitGroup

func increment(from string) {
	for i := 0; i < 10; i++ {
		x := counter
		x++
		time.Sleep(time.Second)
		counter = x
		fmt.Println(from, ":", counter)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go increment("Red")
	go increment("Blue")
	wg.Wait()
	fmt.Println("Final count value is", counter)
}

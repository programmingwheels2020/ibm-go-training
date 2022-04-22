package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var wg sync.WaitGroup
var mutex sync.Mutex

func increment(from string) {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		x := counter
		x++
		time.Sleep(time.Second)
		counter = x
		fmt.Println(from, ":", counter)
		mutex.Unlock()
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

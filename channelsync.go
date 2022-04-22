package main

import "fmt"

func myFunction(from string, done chan bool) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
	done <- true
}

func main() {
	done := make(chan bool)
	go myFunction("Red", done)
	<-done
}

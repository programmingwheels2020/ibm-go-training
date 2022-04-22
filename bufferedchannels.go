package main

import "fmt"

func main() {
	messages := make(chan string, 3)
	messages <- "One"
	messages <- "Two"
	messages <- "Three"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {

		fmt.Println("This is from go routine")
		messages <- "Ping"
		messages <- "Pong"
	}()

	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)

}

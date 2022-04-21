package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number ")
	input, _ := reader.ReadString('\n')
	number, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 32)
	if number%2 == 0 {
		fmt.Println("Entered number is even")
	} else {
		fmt.Println("Entered number is odd")
	}
}

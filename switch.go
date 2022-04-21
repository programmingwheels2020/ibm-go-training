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
	fmt.Println(number)
	switch number {
	case 1:
		fmt.Println("MONDAY")
	case 2:
		fmt.Println("TUESDAY")
	case 3:
		fmt.Println("WEDNESDAY")
	case 4:
		fmt.Println("THURSDAY")
	case 5:
		fmt.Println("FRIDAY")
	case 6:
		fmt.Println("SATURDAY")
	case 7:
		fmt.Println("SUNDAY")
	default:
		fmt.Println("Invalid Number entered")
	}
}

package main

import (
	"errors"
	"fmt"
)

func sum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("Negative values are not allowed")
	} else {
		return a + b, nil
	}
}

func main() {
	total, err := sum(10, 10)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(total)
	}

}

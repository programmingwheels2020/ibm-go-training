package main

import "fmt"

type rectangle struct {
	height, width int
}

func (r rectangle) perimeter() int {
	return 2 * (r.height + r.width)
}

func (r *rectangle) area() int {
	return r.height * r.width
}

func newRectagle(height, width int) rectangle {
	return rectangle{height: height, width: width}
}

func main() {
	r1 := newRectagle(200, 300)
	fmt.Println(r1.perimeter())
	r2 := &r1
	fmt.Println(r2.area())
}

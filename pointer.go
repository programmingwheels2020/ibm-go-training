package main

import "fmt"

func changeByVal(myVal int) {
	myVal = 0
}

func changeByRef(myPtr *int) {
	*myPtr = 0
}

func main() {
	var i int = 10
	/*fmt.Println("Address of i is", &i)

	var ptr *int
	ptr = &i
	fmt.Println(*ptr)*/
	//------
	changeByVal(i)
	fmt.Println(i)
	changeByRef(&i)
	fmt.Println(i)

}

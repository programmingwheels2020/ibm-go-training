package main

import "fmt"

func main() {
	var myMap = map[string]int{"MON": 1, "TUE": 2, "WED": 3}
	myMap["THUR"] = 4
	myMap["FRI"] = 5
	fmt.Println(myMap)
	delete(myMap, "FRI")
	fmt.Println(myMap)
	var myAnotherMap = make(map[string]int)
	myAnotherMap["Red"] = 7
	myAnotherMap["Violet"] = 1
	myAnotherMap["Pink"] = 0
	val, ok := myAnotherMap["Blue"]

	fmt.Println("Value is", val, "and ok is ", ok)

	val, ok = myAnotherMap["Pink"]

	fmt.Println("Value is", val, "and ok is ", ok)
}

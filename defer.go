package main

import "fmt"

func connectToDb() {
	fmt.Println("Connecting to DB")
}

func closeDbConnection() {
	fmt.Println("Closing Db Connection")
}

func closeFiles() {
	fmt.Println("This will close file connection")
}

func main() {
	connectToDb()
	defer closeDbConnection()
	defer closeFiles()
	fmt.Println("This is db operation")
	fmt.Println("This is db operation2")
	fmt.Println("This is db operation3")
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func createFile() {
	fmt.Println("Creating a file and writing to it")

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatalf("Failed to create file %s", err)
	}
	defer file.Close()

	_, err = file.WriteString("Lenin Lawrence")
	if err != nil {
		log.Fatalf("Failed to write file %s", err)
	}
	fmt.Printf("\nFile name is %s", file.Name())

}

func readFile() {
	fmt.Println("Read File")
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatalf("Failed to Read file %s", err)
	}
	fmt.Println(string(data))

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}

func main() {
	createFile()
	readFile()
}

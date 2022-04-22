package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HelloWorldRouter)
	http.ListenAndServe(":3000", nil)
}

func HelloWorldRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

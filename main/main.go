package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greeting(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetingHandler(w http.ResponseWriter, r *http.Request) {
	Greeting(w, "world")
}

// これは動かない
func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetingHandler))
}

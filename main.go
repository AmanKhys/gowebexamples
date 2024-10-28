package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	}))

	err := http.ListenAndServe(":8800", nil)
	if err != nil {
		fmt.Println(err)
	}
}

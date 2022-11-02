package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello aws"))
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Printf("%s", err)
	}
}
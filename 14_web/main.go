package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h2> Hello World </h2>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Starting Server...")
	http.ListenAndServe(":4000", nil)

}

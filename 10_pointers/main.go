package main

import "fmt"

func main() {
	// pointers are used to get the memory address of a value

	a := 5
	b := &a

	fmt.Printf("%T\n", b)
	fmt.Println(a, b)
}

package main

import "fmt"

func main() {

	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Shorter Method

	for i := 0; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}
}

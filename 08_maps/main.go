package main

import (
	"fmt"
)

func main() {
	// Create a map Assigning key values to map
	emails := map[string]string{"John": "john@gmail.com", "teddy": "teddy@gmail.com", "nonso": "nonsoamadi@aol.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))

	// Assigning new Values
	emails["bryan"] = "bryan@yahoo.com"
	fmt.Println(emails)
}

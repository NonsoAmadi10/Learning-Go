package main

import "fmt"

func main() {

	// Arrays - arrays in go are of fixed length

	// var hobbiesArr [3]string

	// // Assign values

	// hobbiesArr[0] = "Swimming"
	// hobbiesArr[1] = "Coding"
	// hobbiesArr[2] = "Sleeping"

	// A shorter way
	// hobbiesArr := [3]string{"Coding", "Swimming", "Sleeping"}

	// fmt.Println(hobbiesArr)

	//Slices - slices are arrays of varying length

	hobbiesSlice := []string{"coding", "swimming", "sleeping"}
	fmt.Println(hobbiesSlice)
	fmt.Println(hobbiesSlice[0:2])
}

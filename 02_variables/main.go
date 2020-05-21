package main

import "fmt"

func main() {

	//Using Var

	var name string
	name = "Nonso Amadi"

	var age int = 23

	//Using const

	const isCool = true
	//isCool = false - cannot reassign variable when using const

	//Using a shorthand method
	isMale := true //This cannot be declared outside a function body

	fmt.Println(name, age, isCool, isMale)

	// TO get the type of a variable

	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", name)
}

package main

import (
	"fmt"
	"strconv"
)

// Person structs
type Person struct {
	firstName string
	lastName  string
	gender    string
	age       int
}

// There are two types of method in structs - value receivers and pointer methods

//an example of value receivers

func (p Person) greet() string {
	return "Hello " + "my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years"

} // An example of pointer method

func (p *Person) hasBirthday() {
	p.age++
}

func main() {

	person1 := Person{"Amadi", "Nonso", "m", 23}

	fmt.Println(person1.age)

	person1.hasBirthday()

	fmt.Println(person1.greet())

}

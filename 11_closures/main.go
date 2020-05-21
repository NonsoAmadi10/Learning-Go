package main

import "fmt"

// Go has  a way of declaring anonymous function i.e functions without a name

func adder() func(int) int {
	sum := 0

	return func(x int) int { // anonymous function
		sum += x

		return sum
	}
}

func main() {
	sum := adder()

	for i := 0; i <= 10; i++ {
		fmt.Println(sum(i))
	}

}

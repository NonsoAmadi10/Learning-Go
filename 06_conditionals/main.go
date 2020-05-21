package main

import "fmt"

func main() {

	x := 7
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d \n", x, y)
	} else {
		fmt.Printf("%d is less than %d \n", y, x)
	}

	//if-elseif-else
	itRains := true
	itShines := false

	if itRains && !itShines {
		fmt.Println("Its Raining outside")
	} else if !itRains && itShines {
		fmt.Println("Its a sunny day today")
	} else {
		fmt.Println("Its a funny day today mate")
	}

	//Switch

	favColor := ""

	switch favColor {
	case "Grey":
		fmt.Println("Your favorite color is grey")
	case "Red":
		fmt.Println("Red signifies Danger")

	default:
		fmt.Println("I have no idea")
	}

}

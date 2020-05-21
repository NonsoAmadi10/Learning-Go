package main

import "fmt"

func main() {
	//range is used to loop over maps, slices and arrays

	idSlice := []int{12, 45, 76, 88, 90}

	for i, id := range idSlice {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range idSlice {
		fmt.Printf("ID - %d \n", id)
	}

	sum := 0

	for _, id := range idSlice {
		sum += id
	}
	fmt.Printf("Sum - %d\n", sum)

	//Range with map

	myProfile := map[string]string{"Name": "Amadi Justice", "Profession": "Software Engineer", "favFood": "Vegetables"}

	for k, v := range myProfile {
		fmt.Printf("%s - %s \n", k, v)
	}
}

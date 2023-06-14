package main

import "fmt"

func main() {
	arr := []int{}

	for i := 1; i < 12; i++ {
		arr = append(arr, i)
	}
	fmt.Printf("slice values: %v \n", arr)

	for _, i := range arr {

		if i%2 == 0 {
			fmt.Printf("%v is even \n", i)
		} else {
			fmt.Printf("%v is odd \n", i)
		}
	}
	for _, i := range arr {
		fmt.Printf("%v is Even: %t \n", i, isEven(i))

	}

}

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	} else {
		return false
	}
}

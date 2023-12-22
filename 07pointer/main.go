package main

import "fmt"

func main() {
	fmt.Println("Welcome a class on pointers")

	// var ptr *int
	// fmt.Println("value og pointer is: ", ptr)

	myNumber := 23

	var pter = &myNumber
	fmt.Println("value of actual pointer is ", *pter)

	*pter = *pter + 2
	fmt.Println("New value is: ", myNumber)
	fmt.Println("value: ", *pter)
}

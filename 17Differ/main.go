package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

// defer stores into the stack

// world, one, two

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // reverse print
		//4 3 2 1 0
	}
}

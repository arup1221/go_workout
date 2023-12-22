package main

import "fmt"

func main() {
	fmt.Println("welcome to funtions in golang")
	greeterTwo()

	proRes, myMessage := proadder(2, 3, 5, 7)
	
	fmt.Println(proRes)
	fmt.Println(myMessage)
}

// here you can not write a function inside a function
func greeterTwo() {
	fmt.Println("Another method")
	greeter()

	zresult := adder(3, 5)
	fmt.Println("Result is: ", zresult)
}

func greeter() {
	fmt.Println("Namstay from golang")
}
func adder(valone int, valtwo int) int {
	return valone + valtwo
}

func proadder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	// return total

	return total, "hi hello"
}

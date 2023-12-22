package main

import "fmt"

func main() {
	fmt.Println("If Else in golang")

	loginCount := 2
	var result string

	if loginCount == 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Something else"
	} else {
		result = "Error"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num := 3; num < 10 {
		fmt.Println("less then 10")
	} else {
		fmt.Println("Not less than 10")
	}

	// if err != nill {

	//}

}

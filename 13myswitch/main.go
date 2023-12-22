package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Swich case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1")

	case 2:
		fmt.Println("dice value is 2 ")
		fallthrough //go to the next statement just like not use of break in java
	case 3:
		fmt.Println("dice value is 3")
		fallthrough
	case 4:
		fmt.Println("dice value is 4")

	case 5:
		fmt.Println("dice value is 5")

	case 6:
		fmt.Println("dice value is 6")

	default:
		fmt.Println("default")
	}
}

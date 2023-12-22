package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to math in golang")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4

	// fmt.Println("The sum is: ", mynumberOne+int(mynumberTwo))   // bad syntax

	//random number
	// rand.Seed(time.Now().UnixNano()) // deprecated
	// fmt.Println(rand.Intn(5) + 1)

	//random from crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "orange"
	fruitList[2] = "coconut"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list Lenght is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mashroom"}
	fmt.Println("vegy List is: ", vegList)
	fmt.Println("vegy list length is : ", len(vegList))
}

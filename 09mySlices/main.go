package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices video")

	var fruitList = []string{"apple", "Orange", "Coconut"}
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana") // this is the method to add elements in a array
	fmt.Println(fruitList)
	// fruitList = append(fruitList[1:]) // delete first element
	// fmt.Println(fruitList)

	fruitList = append(fruitList[1:4])
	fmt.Println(fruitList)

	highScores := make([]int, 4) //slices

	highScores[0] = 239
	highScores[1] = 247
	highScores[2] = 210
	highScores[3] = 209
	//highScores[4] = 274 // but this gives index out of bound

	fmt.Println(highScores)

	highScores = append(highScores, 456, 36) // realocate the memory
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores) // for sorting asending order
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value form slices based on index

	var courses = []string{"reactJS", "JavaScript","Python", "golang", "Tailwind", "Rust"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	days := []string{"Sunday", "Tuesday", "Friday", "Saturday"}

	// fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	roughValue := 1
	for roughValue < 10 {

		if roughValue == 2 {
			goto lco
		}

		if roughValue == 5 {
			break
		}
		fmt.Println("Value is: ", roughValue)
		roughValue++

	}
lco:
	{
		fmt.Println("hey Hitesh")
	}

}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps is golang")

	languages := make(map[string]string) //map[key]string

	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["py"] = "python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS stands for : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang   [:=  walrus operator]
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}

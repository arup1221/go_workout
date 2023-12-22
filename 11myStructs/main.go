package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// no inheritance in golang // no super or parent

	arup := User{"Arup", "gopegmai.com", true, 21}
	fmt.Println(arup)
	fmt.Printf("Arup details are: %+v\n", arup)
	fmt.Printf("Name is %v, Age is %v", arup.Name, arup.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

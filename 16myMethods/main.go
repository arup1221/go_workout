package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// no inheritance in golang // no super or parent

	arup := User{"Arup", "gopegmai.com", true, 21}
	fmt.Println(arup)
	fmt.Printf("Arup details are: %+v\n", arup)
	fmt.Printf("Name is %v, Age is %v", arup.Name, arup.Age)
	arup.GetStatus()
	arup.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("\nIs user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "Gsfds@fmai.com"
	fmt.Println("Email of this user is:", u.Email)
}

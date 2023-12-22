package main

import "fmt"

// for global declaration ,here you cannot use walrus operator ":="
var jwtToken int = 300000

const LoginToken string = "ghabbhjd" //public

func main() {
	var username string = "arup"
	fmt.Println(username)
	fmt.Printf("veriable is of type :%T\n", username)

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("veriable is of type :%T\n", isLogged)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("veriable is of type :%T\n", smallVal)

	var smallFloat float32 = 255.45553432423
	fmt.Println(smallFloat)
	fmt.Printf("veriable is of type :%T\n", smallFloat)

	//default values and some aliases
	var anotherVeriable int
	fmt.Println(anotherVeriable)
	fmt.Printf("veriable is of type :%T\n", anotherVeriable)

	var otherVeriable string
	fmt.Println(otherVeriable)
	fmt.Printf("veriable is of type :%T\n", otherVeriable)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style  [:= walrus operator]   short variable declaration operator
	numberofUser := 30000.43
	fmt.Println(numberofUser)

	fmt.Println(LoginToken)
	fmt.Printf("veriable is of type :%T\n", LoginToken)
}

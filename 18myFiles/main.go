package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files in Golang")

	content := "This needs to go in a file - LearncodeOnline.in"

	file, err := os.Create("./mylcofile.txt") // to create a file

	checkNillErr(err)
	length, err := io.WriteString(file, content) // to write
	// it gives back length not the file itself

	// if err != nil {
	// 	panic(err)
	// }
	checkNillErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcofile.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)

	checkNillErr(err)

	fmt.Println("Text data inside the file is: \n", string(databyte))
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}

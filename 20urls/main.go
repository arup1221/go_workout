package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handeling urls")
	fmt.Println(myUrl)

	//parsing
	result, _ := url.Parse(myUrl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(result.Query())

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Params is: ", val)
	}

	partsofUrl := &url.URL{ //here passing referances not the copy of the url
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutss",
		RawPath: "/user=hitesh",
	}
	anotherUrl := partsofUrl.String()

	fmt.Println(anotherUrl)
}

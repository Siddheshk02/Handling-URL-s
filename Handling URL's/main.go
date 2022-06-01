package main

import (
	"fmt"
	"net/url"
)

var myurl string = "https://lco.dev:2000/learn?coursename=reactjs&paymentid=abc12345ijk"

func main() {
	fmt.Println("Welcome to Handling URL's in Golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	//fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, i := range qparams {
		fmt.Println("Param is: ", i)
	}

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "www.google.com",
		Path:    "/learn",
		RawPath: "user=Sid",
	}

	anotherURL := partsofUrl.String()
	fmt.Println(anotherURL)
}

package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Urls")

	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println("Scheme : ", result.Scheme)
	fmt.Println("Hostname : ", result.Hostname())
	fmt.Println("Path : ", result.Path)
	fmt.Println("Raw Query : ", result.RawQuery)
	fmt.Println("Port Number : ", result.Port())

	qparams := result.Query()
	fmt.Printf("The Type of Query Parameters are : %T", qparams)

	// fmt.Println("Course Name : ", qparams["coursename"])
	// fmt.Println("Payment ID : ", qparams["paymentid"])

	for _, value := range qparams {
		fmt.Println("Param is : ", value)
	}

	partsOfURL := &url.URL{
		Scheme:  "http",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=mihir",
	}

	anotherURL := partsOfURL.String()
	fmt.Println("Another Url : ", anotherURL)
}

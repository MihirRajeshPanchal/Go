package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com/webhp?hl=en&sa=X&ved=0ahUKEwi-lJX4kej1AhU9SmwGHSxmDFYQPAgI"

func main() {
	fmt.Println("Web Request")

	response, err := http.Get(url)

	checkNilError(err)

	//fmt.Println("Response is : ", response)

	fmt.Printf("Response is of Type : %T\n", response)

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	checkNilError(err)

	fmt.Println("Content of URL is : ", string(data))
}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

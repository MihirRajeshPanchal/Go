package main

import (
	"fmt"
)

func main() {
	fmt.Println("If Else")

	var num int

	fmt.Print("Enter a Number : ")
	fmt.Scanln(&num)

	//ifelse
	if num%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}

	//ifelse ladder
	if num > 0 {
		fmt.Println("Number is Positive")
	} else if num < 0 {
		fmt.Println("Number is Negative")
	} else {
		fmt.Println("Number is Zero")
	}
}

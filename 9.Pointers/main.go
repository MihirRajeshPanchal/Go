package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// var ptr *int
	// fmt.Println(ptr)

	myNumber := 18
	var ptr = &myNumber

	fmt.Println("Memory Address of Pointer is : ", ptr)
	fmt.Println("Value stored in address pointed by pointer is : ", *ptr)

	*ptr = *ptr * 2

	fmt.Println("New Value is : ", *ptr)

}

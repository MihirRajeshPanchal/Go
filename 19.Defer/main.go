package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("Hello World")
	fmt.Println("Defer")
	myDiffer()
}

func myDiffer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

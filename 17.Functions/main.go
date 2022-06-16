package main

import "fmt"

func main() {
	fmt.Println("Main")
	greet()

	result := adder(2, 3)
	fmt.Println("Result of Addition is : ", result)

	proResult := proAdder(2, 3, 1, 3, 5, 6, 7, 89, 0)
	fmt.Println("Result of Addition is : ", proResult)

	proMessageResult, message := proMessageAdder(34, 54, 2, 1, 23, 44, 332, 53)
	fmt.Println("Result of Addition is : ", proMessageResult)
	fmt.Println("Pro Message : ", message)
}

func adder(x int, y int) int {
	return x + y
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
func proMessageAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Message from function proMessageAdder()"
}

func greet() {
	fmt.Println("Hello")
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Rating : ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating : ", input)
	fmt.Printf("Type of Rating is : %T", input)

	var num int
	fmt.Print("\nEnter a number : ")
	fmt.Scanln(&num)
	fmt.Println("Number is : ", num)

}

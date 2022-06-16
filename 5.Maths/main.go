package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Maths")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter First Number : ")
	inputOne, _ := reader.ReadString('\n')
	fmt.Print("Enter Second Number : ")
	inputTwo, _ := reader.ReadString('\n')

	numValueOne, err := strconv.ParseFloat(strings.TrimSpace(inputOne), 64)
	numValueTwo, err := strconv.ParseFloat(strings.TrimSpace(inputTwo), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Addition of two Numbers is : ", numValueOne+numValueTwo)
	}

}

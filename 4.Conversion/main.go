package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Conversion")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Rating between 1-5 : ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating : ", input)

	numValue, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Rating + 1 = ", numValue+1)
	}

}

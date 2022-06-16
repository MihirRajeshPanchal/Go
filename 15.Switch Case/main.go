package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice Rolled : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Move 1 space")
		//fallthrough
	case 2:
		fmt.Println("Move 2 space")
	case 3:
		fmt.Println("Move 3 space")
	case 4:
		fmt.Println("Move 4 space")
	case 5:
		fmt.Println("Move 5 space")
	case 6:
		fmt.Println("Move 6 space")

	}
}

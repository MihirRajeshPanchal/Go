package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	//random
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println(myRandomNum)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//random
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(time.Now().Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2020, time.August, 16, 12, 0, 0, 0, time.UTC)

	fmt.Println(createdDate)
}

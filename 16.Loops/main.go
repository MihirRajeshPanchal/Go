package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Day Index : %v \t Day : %v\n", index, day)
	// }

	value := 1

	for value <= 10 {
		if value == 5 {
			value++
			continue
		}
		if value == 8 {
			goto mp
		}
		fmt.Println(value)
		value++
	}

mp:
	fmt.Println("Goto Label")
}
